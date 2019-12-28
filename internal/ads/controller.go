package ads

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

// Video type to hold information about the video
type Video struct {
	ID string `json:"id"`
}

// VidoCtx contains the request information for the requested Video ID
func VidoCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		video := Video{ID: id}
		ctx := context.WithValue(r.Context(), "video", video)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Controller struct to hold our routes
type Controller struct{}

// NewController creates a new Ads controller
func NewController() *Controller {
	return &Controller{}
}

// Ad contains information about the individual ads
type Ad struct {
	Start  int `json:"start"`
	Length int `json:"end"`
	Score  int `json:"score"`
}

type getAdResponse struct {
	Video Video `json:"video"`
	Ads   []Ad  `json:"ads"`
}

// Return the ads for a single video ID
func (c Controller) handleGetAd(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	video, ok := ctx.Value("video").(Video)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	var ads []Ad

	// Actually get these out of a DB
	ads = append(ads, Ad{0, 10, 5})
	ads = append(ads, Ad{0, 7, -3})
	ads = append(ads, Ad{35, 25, 5})

	resp := getAdResponse{
		Video: video,
		Ads:   ads,
	}
	json.NewEncoder(w).Encode(resp)
}
