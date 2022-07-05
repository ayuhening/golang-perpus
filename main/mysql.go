package main

import (
	"database/sql"
	"log"
)

func database() *sql.DB {
	db, err := sql.Open("mysql", "root:password01@tcp(127.0.0.1:3306)/perpus")

	if err != nil {
		log.Fatal(err)
	}
	return db

}
