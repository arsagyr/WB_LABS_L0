package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	"wb_labs_l0/backend/internal/database"
	"wb_labs_l0/backend/internal/model"

	_ "github.com/lib/pq"
)

func OrderTablePage(w http.ResponseWriter, r *http.Request) {
	order := database.GetOrderByID("1")
	orders := []model.Order{order}
	CreateOrderJSON(orders)
	GetOrderDBTable(w, r, orders)
}
func AllOrdersTablePage(w http.ResponseWriter, r *http.Request) {
	orders := database.GetOrders()
	CreateOrderJSON(orders)
	GetOrderDBTable(w, r, orders)
}
func CreateOrderJSON(orders []model.Order) {
	data, err := json.MarshalIndent(orders, "", "  ")
	if err != nil {
		log.Println(err)
	}

	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	d := time.Now().Sub(start)
	s := d.String()
	s = "cache/" + s + "orders.json"
	err = os.WriteFile(s, data, 0644)
	if err != nil {
		log.Println(err)
	}
}

func GetOrderDBTable(w http.ResponseWriter, r *http.Request, orders []model.Order) {
	tmpl, _ := template.ParseFiles("frontend/templates/orders.html")
	tmpl.Execute(w, orders)
}
