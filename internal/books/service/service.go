package service

import (
	books "book-tracker/internal/books/models"
	rep "book-tracker/internal/books/repository"
	"database/sql"
	"log"
)

func GetBooks(db *sql.DB) (*sql.Rows, error) {
	b, err := rep.GetBooks(db)
	if err != nil {
		log.Printf("Error getting books: %v", err)
		return nil, err
	}

	return b, nil
}

func CreateBook(db *sql.DB, b *books.Book) (sql.Result, error) {
	res, err := rep.InsertBook(db, b)
	if err != nil {
		log.Printf("Error inserting book: %v", err)
		return nil, err
	}

	return res, nil
}