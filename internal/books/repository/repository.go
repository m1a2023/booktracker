/*
*
books.repository package contains funcs,
querying database
*/
package rep

import (
	books "book-tracker/internal/books/models"
	"database/sql"
	"encoding/json"
)

func GetBooks(db *sql.DB) (*sql.Rows, error) {
	const st = `SELECT * FROM books;`
	return db.Query(st)
}

func InsertBook(tx *sql.Tx, b *books.Book) (*sql.Rows, error) {
	const st = `INSERT INTO books (title, authors, rate) 
		VALUES (?, ?, ?) RETURNING id;`

	authors, err := json.Marshal(b.Authors)
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(st, b.Title, authors, b.Rate)
	if err != nil {
		return nil, err
	}	
	
	return rows, nil
}