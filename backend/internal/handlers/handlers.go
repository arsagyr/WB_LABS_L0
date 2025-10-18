package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	model "wb_labs_l0/backend/internal/model"
)

var DB *sql.DB

func GetActors(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	rows, err := DB.Query(`
	
	`)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	actors := []model.Actor{}

	for rows.Next() {
		a := model.Actor{}
		err := rows.Scan(&a.Id, &a.Familyname, &a.Givenname, &a.Nation, &a.Number, &a.Honorar)
		if err != nil {
			fmt.Println(err)
			continue
		}
		actors = append(actors, a)
	}

	json.NewEncoder(w).Encode(actors)
}
