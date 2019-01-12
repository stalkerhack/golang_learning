package main

import (
	"database/sql"
	//"fmt"
	"log"

	//"github.com/gotk3/gotk3/gtk"

	_ "github.com/lib/pq"
)

type PC struct {
	ID      string
	Inv_num int
	Name    string
}

func main() {
	db, err := sql.Open("postgres", "postgres://login:password123@localhost/books_shop")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}
