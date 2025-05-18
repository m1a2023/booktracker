/*
*
book.service package contains funcs,
calling repository and transforming data
*/
package service

import (
	books "book-tracker/internal/books/models"
	rep "book-tracker/internal/books/repository"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func GetBooks(db *sql.DB) ([]*books.Book, error) {
	brows, err := rep.GetBooks(db)
	if err != nil {
		log.Printf("Error getting books: %v", err)
		return nil, err
	}

	defer brows.Close()

	data := []*books.Book{}
	for brows.Next() {
		b := &books.Book{}

		var authors string
		if err := brows.Scan(&b.Id, &b.Title, &authors, &b.Rate); err != nil {
			msg := fmt.Sprintf("Error while scanning: %v", err)
			fmt.Println(msg)
			continue
		}

		b.Authors = strings.Split(authors, ",")
		data = append(data, b)
	}

	return data, nil
}

func CreateBook(db *sql.DB, b *books.Book) (int32, error) {
	id, err := rep.InsertBook(db, b)
	if err != nil {
		log.Printf("Error inserting book: %v", err)
		return -1, err
	}

	return id, nil
}