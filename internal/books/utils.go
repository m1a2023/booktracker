package utils

import (
	books "book-tracker/internal/books/models"
	rep "book-tracker/internal/books/repository"
	"database/sql"
	"encoding/json"
	"net/http"
)

func MultipleRowsBuildResponse(r *sql.Rows) ([]*books.Book, error) {
	// Building response 
	res := []*books.Book{}
	for r.Next() {
		b := &books.Book{}

		var authors string
		if err := r.Scan(&b.Id, &b.Title, &authors, &b.Rate); err != nil {
			continue
		}

		err := json.Unmarshal([]byte(authors), &b.Authors)
		if err != nil {
			return nil, err
		}
		res = append(res, b)
	}

	return res, nil
}

func TryCreateBook(con *sql.DB, b *books.Book) (int, bool) {
	// Book's saved id 
	var id int 

	// Start a transaction 
	tx, err := con.Begin()
	if err != nil {
		return 0, false  
	}

	// Insert book
	rows, err := rep.InsertBook(tx, b)
	if err != nil {
		tx.Rollback()
		return 0, false  
	}

	// Gets book's id
	rows.Next() 

	if err := rows.Scan(&id); err != nil {
		tx.Rollback()
		return 0, false
	}

	// Commit changes
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return 0, false
	}

	return id, true
}


type SingleIdResponse struct { 
	Id 		int 		`json:"id"` 
}


func SetGetHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func SetPostHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}


func SetUpdateHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}


func SetDeleteHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}