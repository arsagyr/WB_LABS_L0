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
	CreateOrderDBTable()
	GetOrderDBTable(w, r)
}
func CreateOrderDBTable() {
	orders := database.GetOrdersDB()

	data, err := json.MarshalIndent(orders, "", "  ")
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile("frontend/json/table.json", data, 0644)
	if err != nil {
		log.Println(err)
	}
}

func GetOrderDBTable(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("frontend/json/table.json")
	if err != nil {
		log.Println(err)
	}

	var orders []model.Order
	err = json.Unmarshal(file, &orders)
	if err != nil {
		panic(err)
	}

	tmpl, _ := template.ParseFiles("frontend/templates/orders.html")
	tmpl.Execute(w, orders)
}
