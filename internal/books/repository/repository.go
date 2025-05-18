/*
*
books.repository package contains funcs,
querying database
*/
package rep

import (
	books "book-tracker/internal/books/models"
	"database/sql"
	"log"
	"strings"
	"sync"
)

func GetBooks(db *sql.DB) (*sql.Rows, error) {
	st := `SELECT * FROM books;`
	return db.Query(st)
}

var mut sync.Mutex

func InsertBook(db *sql.DB, b *books.Book) (int32, error) {
	st := `INSERT INTO books (title, authors, rate) 
		VALUES (?, ?, ?) RETURNING id;`

	mut.Lock()
	defer mut.Unlock()

	insertBook, err := db.Prepare(st)
	if err != nil {
		log.Printf("Prepare statement incorrect: %v\n", err)
		return -1, err
	}

	defer insertBook.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Printf("Transaction fault: %v\n", err)
		return -1, err
	}

	var id int32
	err = tx.Stmt(insertBook).QueryRow(b.Title, strings.Join(b.Authors, ","), b.Rate).Scan(&id)
	if err != nil {
		tx.Rollback()
		log.Printf("Insertion failed: %v\n", err)
		return -1, err
	}
	
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		log.Printf("Transaction commit failed: %v", err)
		return -1, err
	}

	return id, nil
}