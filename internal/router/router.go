package router

import (
	"net/http"

	"github.com/NikhilSharma03/Mirai/internal/app"

	"github.com/go-chi/chi/v5"
)

func NewRouter(app *app.App) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte("welcome"))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	return r
}
