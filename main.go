package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/danieldides/skipper-server/internal/ads"
	"github.com/danieldides/skipper-server/internal/util"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"net/http"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	const port = ":8080"

	r := chi.NewRouter()

	dbCfg := util.DBOpts{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}

	db, err := util.ConnectDB(dbCfg)
	if err != nil {
		log.Fatal(err)
	}

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// Healthcheck endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		type Response struct {
			Status  int    `json:"status_code"`
			Message string `json:"message"`
		}
		json.NewEncoder(w).Encode(Response{Status: http.StatusOK, Message: "It's working"})
	})

	r.Group(func(r chi.Router) {
		r.Mount("/ads", ads.New(db))
	})

	log.Printf("Serving HTTP on %v\n", port)
	http.ListenAndServe(port, r)
}
