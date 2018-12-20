package main

import {"fmt"
		"database/sql"
		_ "github.com/lib/pq"
}

type books struct {
	isbn string
	title string
	author string
	price float32	
}

func main() {
		
}