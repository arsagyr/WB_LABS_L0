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

func PaymentTablePage(w http.ResponseWriter, r *http.Request) {
	payments := database.GetPaymentByID(1)
	CreatePaymentJSON(payments)
	GetPaymentDBTable(w, r, payments)
}
func AllPaymentsTablePage(w http.ResponseWriter, r *http.Request) {
	payments := database.GetAllPayments()
	CreatePaymentJSON(payments)
	GetPaymentDBTable(w, r, payments)
}
func CreatePaymentJSON(payments []model.Payment) {
	data, err := json.MarshalIndent(payments, "", "  ")
	if err != nil {
		log.Println(err)
	}
	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	d := time.Now().Sub(start)
	s := d.String()
	s = "cache/" + s + "payments.json"
	err = os.WriteFile(s, data, 0644)
	if err != nil {
		log.Println(err)
	}
}

func GetPaymentDBTable(w http.ResponseWriter, r *http.Request, payments []model.Payment) {
	tmpl, _ := template.ParseFiles("frontend/templates/payments.html")
	tmpl.Execute(w, payments)
}
