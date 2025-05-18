package handlers

import (
	books "book-tracker/internal/books/models"
	service "book-tracker/internal/books/service"
	"book-tracker/internal/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	conn := db.GetConnection()
	books, err := service.GetBooks(conn)
	if err != nil {
		msg := fmt.Sprintf("Error: %v", err) // "Internal server error"
		http.Error(w, msg, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func PostBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	var book *books.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return 
	}

	conn := db.GetConnection()
	id, err := service.CreateBook(conn, book)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	res := struct { 
		Id int32 	`json:"id"` 
	} { Id: id }
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error parsing data to json", http.StatusInternalServerError)
	}
}