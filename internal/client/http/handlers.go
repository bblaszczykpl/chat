package http

import (
	"chat/internal/client/app"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

type HttpHandler struct {
	app *app.Application
}

func NewHttpHandler(app *app.Application) *HttpHandler {
	return &HttpHandler{
		app: app,
	}
}

func (h *HttpHandler) CreateNewRoom(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New room")
}

func (h *HttpHandler) Home(w http.ResponseWriter, r *http.Request) {
	Render(w, "chat", "index", true, &TemplateData{})
}

func (h *HttpHandler) FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
