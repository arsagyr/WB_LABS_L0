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

func DeliveryByIDTablePage(w http.ResponseWriter, r *http.Request) {
	deliveries := []model.Delivery{database.GetDeliveryByID(1)}
	CreateDeliveriesJSON(deliveries)
	GetDeliveries(w, r, deliveries)
}
func AllDeliveriesTablePage(w http.ResponseWriter, r *http.Request) {
	deliveries := database.GetAllDeliveries()
	CreateDeliveriesJSON(deliveries)
	GetDeliveries(w, r, deliveries)
}

func CreateDeliveriesJSON(deliveries []model.Delivery) {
	data, err := json.MarshalIndent(deliveries, "", "  ")
	if err != nil {
		log.Println(err)
	}
	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	d := time.Now().Sub(start)
	s := d.String()
	s = "cache/" + s + "deliveriesall.json"
	err = os.WriteFile(s, data, 0644)
	if err != nil {
		log.Println(err)
	}
}

func GetDeliveries(w http.ResponseWriter, r *http.Request, deliveries []model.Delivery) {
	tmpl, _ := template.ParseFiles("frontend/templates/deliveries.html")
	tmpl.Execute(w, deliveries)
}
