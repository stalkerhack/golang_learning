package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"

	_ "github.com/lib/pq"
)

type PC struct {
	ID      string
	Inv_num int
	Name    string
}

func main() {
	db, err := sql.Open("postgres", "postgres://stalkerhack:4273@localhost/PCs")
	if err != nil {
		log.Fatal(err)
	}
}
