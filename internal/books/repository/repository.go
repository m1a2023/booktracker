package rep

import (
	books "book-tracker/internal/books/models"
	"database/sql"
	"log"
	"strings"
)

func GetBooks(db *sql.DB) (*sql.Rows, error) {
	st := `SELECT * FROM books;`
	return db.Query(st)
}

func InsertBook(db *sql.DB, b *books.Book) (sql.Result, error) {
	st := `INSERT INTO books (title, authors, rate) 
		VALUES (?, ?, ?);`

	insertBook, err := db.Prepare(st)
	if err != nil {
		log.Printf("Prepare statement incorrect: %v\n", err)
		return nil, err
	}

	defer insertBook.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Printf("Transaction fault: %v\n", err)
		return nil, err
	}

	res, err := tx.Stmt(insertBook).Exec(b.Title, strings.Join(b.Authors, ","), b.Rate)
	if err != nil {
		tx.Rollback()
		log.Printf("Insertion failed: %v\n", err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Transaction commit failed: %v", err)
		return nil, err
	}

	return res, nil
}