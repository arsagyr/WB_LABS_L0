package main

import (
	"database/sql"
	"log"
	"wb_labs_l0/backend/internal/database"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=password dbname=orders sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	database.DB = db
	database.CreateOrderFile(1)
	db.Close()
}
