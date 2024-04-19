package http

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var wsChan = make(chan WsPayload)

var clients = make(map[WebSocketConnection]string)

type WebSocketConnection struct {
	*websocket.Conn
}

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type WsJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

type WsPayload struct {
	Action  string              `json:"action"`
	Message string              `json:"message"`
	Conn    WebSocketConnection `json:"-"`
}

func (h *HttpHandler) WsConnection(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected!")

	var response WsJsonResponse
	response.Message = "Hello!"
	ws.WriteJSON(response)

	conn := WebSocketConnection{Conn: ws}
	clients[conn] = ""

	go ListenForWs(&conn)
}

func ListenToWsChannel() {
	log.Println("Starting channel listener...")

	var response WsJsonResponse

	for {
		e := <-wsChan
		response.Message = e.Message
		broadcastToAll(response, e.Conn)
	}

}

func broadcastToAll(response WsJsonResponse, conn WebSocketConnection) {
	for client := range clients {
		if client == conn {
			continue
		}
		err := client.WriteJSON(response)
		if err != nil {
			log.Println("Broadcast error", err)
		}
	}

}

func ListenForWs(conn *WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error:", fmt.Sprintf("%v", r))
		}
	}()

	var payload WsPayload

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {

		} else {
			payload.Conn = *conn
			wsChan <- payload
		}

	}
}
