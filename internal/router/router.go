package router

import (
	"net/http"

	"github.com/NikhilSharma03/Mirai/internal/app"

	"github.com/go-chi/chi/v5"
)

func NewRouter(app *app.App) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("welcome"))
	})

	return r
}
