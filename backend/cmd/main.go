package main

import (
	"database/sql"
	"log"
	"wb_labs_l0/backend/internal/database"

	_ "github.com/lib/pq"
)

// func runServer() {
// 	db, err := sql.Open("postgres", "user=postgres password=password dbname=testactors sslmode=disable")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	database.DB = db
// 	defer database.DB.Close()

// 	router := mux.NewRouter()
// 	router.HandleFunc("/", handlers.IndexHandler)
// 	router.HandleFunc("/create", handlers.CreateHandler)
// 	router.HandleFunc("/edit/{id:[0-9]+}", handlers.EditPage).Methods("GET")
// 	router.HandleFunc("/edit/{id:[0-9]+}", handlers.EditHandler).Methods("POST")
// 	router.HandleFunc("/delete/{id:[0-9]+}", handlers.DeleteHandler)

// 	http.Handle("/", router)
// 	fmt.Println("Server is listening...")

// 	http.ListenAndServe("localhost:8181", nil)
// }

func main() {
	// runServer()
	db, err := sql.Open("postgres", "user=postgres password=password dbname=orders sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	database.DB = db
	database.CreateOrderFile("1")
	db.Close()

}
