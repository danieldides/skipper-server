package ads

import (
	"net/http"

	"github.com/go-chi/chi"
)

// New creates a new Router and
func New() http.Handler {
	r := chi.NewRouter()
	c := NewController()

	r.Route("/", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})
	})

	r.Route("/{id}", func(r chi.Router) {
		r.Use(VidoCtx)
		r.Get("/", c.handleGetAd)
	})

	return r
}
