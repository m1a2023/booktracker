package main

import (
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

	// Http handlers
	log.Fatal(http.ListenAndServe(":1010", nil))
}	
