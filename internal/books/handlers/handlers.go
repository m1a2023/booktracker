package handlers

import (
	utils "book-tracker/internal/books"
	service "book-tracker/internal/books/service"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	// Set headers 
	utils.SetGetHeaders(w)

	// Control flow transfer 
	service.GetBooks(w, r)
}

func PostBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	utils.SetPostHeaders(w)

	// Control flow transfer 
	service.CreateBook(w, r)
}