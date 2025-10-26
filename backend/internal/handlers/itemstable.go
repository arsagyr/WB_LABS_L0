package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"wb_labs_l0/backend/internal/database"
	"wb_labs_l0/backend/internal/model"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func ItemsByIDTablePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	items := database.GetItemsByOrderStringID(id)
	s := CacheFileNameMaker() + "itemsbyid.json"
	CreateItemJSON(items, s)
	GetItemJSON(w, r, s)
}
func AllItemsTablePage(w http.ResponseWriter, r *http.Request) {
	s := CacheFileNameMaker() + "itemsbyid.json"
	CreateItemJSON(database.GetItems(), s)
	GetItemJSON(w, r, s)
}
func CreateItemJSON(items []model.Item, filename string) {
	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		log.Println(err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Println(err)
	}
}

func GetItemJSON(w http.ResponseWriter, r *http.Request, path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	var items []model.Item
	err = json.Unmarshal(file, &items)
	if err != nil {
		panic(err)
	}
	tmpl, _ := template.ParseFiles("frontend/templates/items.html")
	tmpl.Execute(w, items)
}
