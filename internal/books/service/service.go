/*
*
book.service package contains funcs,
calling repository and transforming data
*/
package service

import (
	utils "book-tracker/internal/books"
	books "book-tracker/internal/books/models"
	rep "book-tracker/internal/books/repository"
	"book-tracker/internal/db"
	"encoding/json"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	// Gets connection 
	con := db.GetConnection()

	// Gets rows 
	brows, err := rep.GetBooks(con)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	// Defer rows close
	defer brows.Close()
	
	// Build response
	res, err := utils.MultipleRowsBuildResponse(brows)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Send response
	json.NewEncoder(w).Encode(res)
}


func CreateBook(w http.ResponseWriter, r *http.Request) {
	// Gets connection 
	con := db.GetConnection()
	
	// Build book
	var book *books.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	// Save book in transaction 
	id, ok := utils.TryCreateBook(con, book)
	if ! ok  {
		http.Error(w, "Transaction fault", http.StatusInternalServerError)
		return 
	}

	// Build response
	res := utils.SingleIdResponse{Id: id}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}