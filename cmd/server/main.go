package main

import (
	"book-tracker/internal/books/handlers"
	"book-tracker/internal/db"
	"log"
	"net/http"
)


func main() {
	// Database connection 
	con := db.GetConnection()
	if err := con.Ping(); err != nil {
		log.Fatalln(err)
	}
	
	defer con.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/books", handlers.GetBooks)
	mux.HandleFunc("/book", handlers.PostBook)

	// Http handlers
	log.Fatal(http.ListenAndServe("0.0.0.0:1010", mux))
}	
