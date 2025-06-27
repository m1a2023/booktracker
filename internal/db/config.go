package db

import (
	books "book-tracker/internal/books/models"
	"database/sql"
	"flag"
	"log"

	_ "modernc.org/sqlite"
)

// db vars
var (
	pool 	*sql.DB
	DSN 	string //"file:demo.db?cache=shared&mode=rwc"
)


func init() {
	// Setup flags 
	setupFlags()

	// Opening database 
	var err error 
	pool, err = sql.Open("sqlite", DSN)
	if err != nil {
		log.Fatalln(err)
	}

	// Creating tables if not exists
	_, err = books.CreateBookTable(pool)
	if err != nil {
		log.Fatalf("Could not create Book table; %v", err)
	}
}

func GetConnection() (*sql.DB) {
	return pool
}

func setupFlags() {
	flag.StringVar(&DSN, "dsn", "store.db", "Path to SQLite database file")
}