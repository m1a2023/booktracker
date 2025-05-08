package db

import (
	books "book-tracker/internal/books/models"
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var pool *sql.DB

func init() {
	var err error 
	dsn := "demo.db"

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