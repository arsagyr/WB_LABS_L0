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

func PaymentTablePage(w http.ResponseWriter, r *http.Request) {
	payments := database.GetPaymentByID(1)
	s := CacheFileNameMaker() + "payments.json"
	CreatePaymentJSON(payments, s)
	GetPaymentDBTable(w, r, s)
}
func AllPaymentsTablePage(w http.ResponseWriter, r *http.Request) {
	payments := database.GetAllPayments()
	s := CacheFileNameMaker() + "payments.json"
	CreatePaymentJSON(payments, s)
	GetPaymentDBTable(w, r, s)
}
func CreatePaymentJSON(payments []model.Payment, filename string) {
	data, err := json.MarshalIndent(payments, "", "  ")
	if err != nil {
		log.Println(err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Println(err)
	}
}

func GetPaymentDBTable(w http.ResponseWriter, r *http.Request, path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	var payments []model.Payment
	err = json.Unmarshal(file, &payments)
	if err != nil {
		panic(err)
	}
	tmpl, _ := template.ParseFiles("backend/templates/payments.html")
	tmpl.Execute(w, payments)
}
