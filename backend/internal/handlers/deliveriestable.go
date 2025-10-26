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

func DeliveryByIDTablePage(w http.ResponseWriter, r *http.Request) {
	deliveries := []model.Delivery{database.GetDeliveryByID(1)}
	s := CacheFileNameMaker() + "deliveriesall.json"
	CreateDeliveriesJSON(deliveries, s)
	GetDeliveries(w, r, s)
}
func AllDeliveriesTablePage(w http.ResponseWriter, r *http.Request) {
	s := CacheFileNameMaker() + "deliveriesall.json"
	CreateDeliveriesJSON(database.GetAllDeliveries(), s)
	GetDeliveries(w, r, s)
}

func CreateDeliveriesJSON(deliveries []model.Delivery, filename string) {
	data, err := json.MarshalIndent(deliveries, "", "  ")
	if err != nil {
		log.Println(err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Println(err)
	}
}

func GetDeliveries(w http.ResponseWriter, r *http.Request, path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	var array []model.Delivery
	err = json.Unmarshal(file, &array)
	if err != nil {
		panic(err)
	}
	tmpl, _ := template.ParseFiles("frontend/templates/deliveries.html")
	tmpl.Execute(w, array)
}
