package http

import (
	"chat/internal/client/app"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func initRoutes(app *app.Application) func(r chi.Router) {
	handlers := NewHttpHandler(app)
	return func(r chi.Router) {
		r.Get("/", handlers.Home)
		r.Post("/createRoom", handlers.CreateNewRoom)
	}
}

func InitStatic(app *app.Application) func(r chi.Router) {
	handlers := NewHttpHandler(app)
	return func(r chi.Router) {
		handlers.FileServer(r, "/assets", http.Dir("./assets"))
	}
}

func BootstrapServer(app *app.Application) {
	r := chi.NewRouter()

	r.Group(
		initRoutes(app),
	)

	r.Group(
		InitStatic(app),
	)
	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:" + strconv.Itoa(app.Settings.Port),
	}

	log.Fatal(s.ListenAndServe())
}
