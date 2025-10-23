package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func RunDB() {
	db, err := sql.Open("postgres", "user=postgres password=password dbname=orders sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	db.Close()
}
