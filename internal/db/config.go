package db

import (
	books "book-tracker/internal/books/models"
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var pool *sql.DB
const dsn string = "demo.db" //"file:demo.db?cache=shared&mode=rwc"

func init() {
	var err error 

	// Opening database 
	pool, err = sql.Open("sqlite", dsn)
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