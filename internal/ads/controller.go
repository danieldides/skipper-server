package ads

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

// Request struct with
type Request struct {
	id string
}

// VideoCtx contains the request information for the requested Video ID
func VideoCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		req := Request{id: id}
		ctx := context.WithValue(r.Context(), "request", req)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Controller struct to hold our routes
type Controller struct {
	db *sql.DB
}

// NewController creates a new Ads controller
func NewController(db *sql.DB) *Controller {
	return &Controller{db: db}
}

type getAdsResponse struct {
	Ads []Ad `json:"ads"`
}

func (c Controller) handleGetAdsByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req, ok := ctx.Value("req").(Request)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
	}

	ads, err := GetAdsByVideoID(c.db, req.id)
	if err != nil {
		w.Write([]byte("y'all screwed up"))
		return
	}

	resp := getAdsResponse{
		Ads: ads,
	}
	json.NewEncoder(w).Encode(resp)
}
