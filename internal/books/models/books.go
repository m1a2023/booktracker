package books

import (
	"database/sql"
)

type Book struct {
	Id 			int 			`json:"id"`
	Title   string   	`json:"name"`
	Authors []string 	`json:"authors"`
	Rate    uint     	`json:"rate"`
}

func CreateBookTable(db *sql.DB) (sql.Result, error) {
	st := `CREATE TABLE IF NOT EXISTS books (
			id 			INT 	AUTOINCREMENT PRIMARY KEY NOT NULL,
			title 	CHAR(50)	NOT NULL,
			authors TEXT NOT NULL,
			rate 		INT);`

	return db.Exec(st)
}