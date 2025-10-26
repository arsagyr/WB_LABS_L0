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

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func ItemsByIDTablePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	items := database.GetItemsByOrderStringID(id)
	CreateItemJSON(items)
	GetItemJSON(w, r, items)
}
func AllItemsTablePage(w http.ResponseWriter, r *http.Request) {
	items := database.GetItems()
	CreateAllItemsJSON()
	GetItemJSON(w, r, items)
}
func CreateItemJSON(items []model.Item) {
	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		log.Println(err)
	}
	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	d := time.Now().Sub(start)
	s := d.String()
	s = "cache/" + s + "itemsbyid.json"
	err = os.WriteFile(s, data, 0644)
	if err != nil {
		log.Println(err)
	}
}

func CreateAllItemsJSON() {
	items := database.GetItemsByOrderID(1)
	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		log.Println(err)
	}
	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	d := time.Now().Sub(start)
	s := d.String()
	s = "cache/" + s + "itemsall.json"
	err = os.WriteFile(s, data, 0644)
	if err != nil {
		log.Println(err)
	}
}

func GetItemJSON(w http.ResponseWriter, r *http.Request, items []model.Item) {
	tmpl, _ := template.ParseFiles("frontend/templates/items.html")
	tmpl.Execute(w, items)
}
