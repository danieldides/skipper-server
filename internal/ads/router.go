package ads

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
)

// New creates a new Router and
func New(db *sql.DB) http.Handler {
	r := chi.NewRouter()
	c := NewController(db)

	r.Route("/", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})
	})

	r.Route("/{id}", func(r chi.Router) {
		r.Use(VideoCtx)
		r.Get("/", c.handleGetAdsByID)
	})

	return r
}
