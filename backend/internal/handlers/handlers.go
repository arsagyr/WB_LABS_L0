package handlers

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"encoding/json"

// 	model "wb_labs_l0/backend/internal/model"

// 	_ "github.com/lib/pq"
// )

// // func IndexHandler(w http.ResponseWriter, r *http.Request) {
// // 	CreateTable()
// // 	GetTable(w, r)
// // }

// func LoadOrder(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Content-Type", "application/json")

// 	rows, err := DB.Query(`

// 	`)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer rows.Close()
// 	orders := []model.Order{}

// 	for rows.Next() {
// 		a := model.Order{}
// 		err := rows.Scan(
// 			&a.Order_uid,
// 			&a.Track_number,
// 			&a.Entry,
// 			&a.Delivery,
// 			&a.Payment,
// 			&a.Items,
// 			&a.Locale,
// 			&a.Internal_signature,
// 			&a.Customer_id,
// 			&a.Delivery_service,
// 			&a.Shardkey,
// 			&a.Sm_id,
// 			&a.Date_created,
// 			&a.Oof_shard,
// 		)
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		orders = append(orders, a)
// 	}

// 	json.NewEncoder(w).Encode(orders)
// }
