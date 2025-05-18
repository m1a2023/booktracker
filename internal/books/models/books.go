package books

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Book struct {
	Id 			int 			`json:"id"`
	Title   string   	`json:"title"`
	Authors []string 	`json:"authors"`
	Rate    uint     	`json:"rate"`
}

func CreateBookTable(db *sql.DB) (sql.Result, error) {
	st := `CREATE TABLE IF NOT EXISTS books (
			id 			INTEGER 	PRIMARY KEY,
			title 	CHAR(50)	NOT NULL,
			authors TEXT NOT 	NULL,
			rate 		INT 			CHECK(rate BETWEEN 0 and 5));`

	return db.Exec(st)
}