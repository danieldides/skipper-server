package util

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// DBOpts contains all necessary information for establishing a DB connection
type DBOpts struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// ConnectDB connects to the master DB and returns the connection
func ConnectDB(opts DBOpts) (*sql.DB, error) {
	uri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		opts.Host, opts.Port, opts.User, opts.Password, opts.Database)

	log.Printf("Connecting to database: %v\n", opts.Host)
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("Connected to database successfully")

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(2)

	return db, nil
}
