package database

import (
	"encoding/json"
	"log"
	"os"
	"wb_labs_l0/backend/internal/model"
)

func GetOrderByID(orderID string) model.Order {
	row := DB.QueryRow(`
	SELECT * FROM Orders
	WHERE id =$1   
	`, orderID)

	orderdb := model.OrderDB{}
	err := row.Scan(
		&orderdb.Id,
		&orderdb.Order_uid,
		&orderdb.Track_number,
		&orderdb.Entry,
		&orderdb.Locale,
		&orderdb.Internal_signature,
		&orderdb.Customer_id,
		&orderdb.Delivery_service,
		&orderdb.Shardkey,
		&orderdb.Sm_id,
		&orderdb.Date_created,
		&orderdb.Oof_shard,
		&orderdb.Delivery_id,
		&orderdb.Payment_id,
	)
	if err != nil {
		log.Println(err)
	}
	items := GetItemsByOrderID(orderID)
	payment := GetPaymentByID(orderdb.Payment_id)
	delivery := GetDeliveryByID(orderdb.Delivery_id)
	order := model.Order{
		Order_uid:          orderdb.Order_uid,
		Track_number:       orderdb.Track_number,
		Entry:              orderdb.Entry,
		Delivery:           delivery,
		Payment:            payment,
		Items:              items,
		Locale:             orderdb.Locale,
		Internal_signature: orderdb.Internal_signature,
		Customer_id:        orderdb.Customer_id,
		Delivery_service:   orderdb.Delivery_service,
		Shardkey:           orderdb.Shardkey,
		Sm_id:              orderdb.Sm_id,
		Date_created:       orderdb.Date_created,
		Oof_shard:          orderdb.Oof_shard,
	}

	return order
}

func CreateOrderFile(id string) {

	order := GetOrderByID(id)
	data, err := json.MarshalIndent(order, "", "  ")
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile("cache/order.json", data, 0644)
	if err != nil {
		log.Println(err)
	}
}
