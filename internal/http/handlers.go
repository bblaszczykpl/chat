package http

import (
	"chat/internal/app"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
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

func (h *HttpHandler) Home(w http.ResponseWriter, r *http.Request) {
	content, _ := template.ParseFiles("./views/layout.html", "./views/partials/header.html", "./views/partials/sidebar.html", "./views/partials/footer.html", "./views/chat/index.html")
	w.Header().Set("Content-Type", "text/html")
	err := content.Execute(w, nil)

	if err != nil {
		log.Println("Error executing chat template template", err)
	}
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
