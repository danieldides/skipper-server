package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/danieldides/skipper-server/internal/ads"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"net/http"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	const port = ":8080"

	r := chi.NewRouter()

	dbCfg := dbOpts{
		host:     os.Getenv("DB_HOST"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		database: os.Getenv("DB_NAME"),
		port:     os.Getenv("DB_PORT"),
	}

	db, err := ConnectDB(dbCfg)
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

type dbOpts struct {
	host     string
	port     string
	user     string
	password string
	database string
}

// ConnectDB connects to the master DB and returns the connection
func ConnectDB(opts dbOpts) (*sql.DB, error) {
	uri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		opts.host, opts.port, opts.user, opts.password, opts.database)

	log.Printf("Connecting to database: %v\n", opts.host)
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to database successfully")

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(2)

	return db, nil
}
