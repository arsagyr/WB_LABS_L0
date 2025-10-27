package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"wb_labs_l0/backend/internal/database"
	"wb_labs_l0/backend/internal/model"

	_ "github.com/lib/pq"
)

func OrderTablePage(w http.ResponseWriter, r *http.Request) {
	order := database.GetOrderByID(1)
	orders := []model.Order{order}
	s := CacheFileNameMaker() + "order.json"
	CreateOrderJSON(orders, s)
	GetOrderDBTable(w, r, s)
}
func AllOrdersTablePage(w http.ResponseWriter, r *http.Request) {
	orders := database.GetOrders()
	s := CacheFileNameMaker() + "orders.json"
	CreateOrderJSON(orders, s)
	GetOrderDBTable(w, r, s)
}
func CreateOrderJSON(orders []model.Order, filename string) {
	data, err := json.MarshalIndent(orders, "", "  ")
	if err != nil {
		log.Println(err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Println(err)
	}
}

func GetOrderDBTable(w http.ResponseWriter, r *http.Request, path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	var orders []model.Order
	err = json.Unmarshal(file, &orders)
	if err != nil {
		panic(err)
	}
	tmpl, _ := template.ParseFiles("backend/templates/orders.html")
	tmpl.Execute(w, orders)
}
