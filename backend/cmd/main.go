package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"

	"wb_labs_l0/backend/internal/database"
	"wb_labs_l0/backend/internal/handlers"
	"wb_labs_l0/backend/internal/subscribes"
)

func runServer() {
	db, err := sql.Open("postgres", "user=postgres password=password dbname=orders sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	database.DB = db
	defer database.DB.Close()

	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}
	log.Println(url)
	nc, _ := nats.Connect(url)
	defer nc.Drain()

	subscribes.SubscribeToOrderJSON(nc)
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Mainpage)
	router.HandleFunc("/searchorder", handlers.SearchOrder)
	router.HandleFunc("/itemsByOrder/{id:[0-9]+}", handlers.ItemsByIDTablePage)
	router.HandleFunc("/all_orders", handlers.AllOrdersTablePage)
	router.HandleFunc("/all_items", handlers.AllItemsTablePage)
	router.HandleFunc("/all_deliveries", handlers.AllDeliveriesTablePage)
	router.HandleFunc("/all_payments", handlers.AllPaymentsTablePage)

	http.Handle("/", router)
	fmt.Println("Server is listening...")
	log.Println("http:///localhost:8181")

	http.ListenAndServe("localhost:8181", nil)
}

func main() {
	runServer()
}
