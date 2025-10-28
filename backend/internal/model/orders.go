package model

import (
	"fmt"
)

type Order struct {
	Order_uid          string   `json:"order_uid" db:"order_uid" `       // Order UID
	Track_number       string   `json:"track_number" db:"track_number" ` // track_number
	Entry              string   `json:"entry" db:"entry" `               //
	Delivery           Delivery `json:"delivery" db:"-" `                //
	Payment            Payment  `json:"payment" db:"-" `                 //
	Items              []Item   `json:"items" db:"-"`
	Locale             string   `json:"locale" db:"locale"`
	Internal_signature string   `json:"internal_signature" db:"internal_signature"`
	Customer_id        string   `json:"customer_id" db:"customer_id"`           //
	Delivery_service   string   `json:"delivery_service" db:"delivery_service"` //
	Shardkey           string   `json:"shardkey" db:"shardkey"`                 //
	Sm_id              int      `json:"sm_id" db:"sm_id"`                       //
	Date_created       string   `json:"date_created" db:"date_created"`         //
	Oof_shard          string   `json:"oof_shard" db:"oof_shard"`               //
}

type OrderDB struct {
	Id                 int    `json:"id" db:"id"`                      // ID
	Order_uid          string `json:"order_uid" db:"order_uid" `       // Order UID
	Track_number       string `json:"track_number" db:"track_number" ` // track_number
	Entry              string `json:"entry" db:"entry" `               //
	Delivery_id        int    `json:"delivery_id" db:"delivery_id" `   //
	Payment_id         int    `json:"payment_id" db:"payment_id" `     //
	Locale             string `json:"locale" db:"locale"`
	Internal_signature string `json:"internal_signature" db:"internal_signature"`
	Customer_id        string `json:"customer_id" db:"customer_id"`           //
	Delivery_service   string `json:"delivery_service" db:"delivery_service"` //
	Shardkey           string `json:"shardkey" db:"shardkey"`                 //
	Sm_id              int    `json:"sm_id" db:"sm_id"`                       //
	Date_created       string `json:"date_created" db:"date_created"`         //
	Oof_shard          string `json:"oof_shard" db:"oof_shard"`               //
}

func (o Order) Print() {
	fmt.Println(o.Order_uid)
	fmt.Println(o.Track_number)
	fmt.Println(o.Entry)
	o.Delivery.Print()
	o.Payment.Print()
	o.Items[0].Print()
	fmt.Println(o.Locale)
	fmt.Println(o.Internal_signature)
	fmt.Println(o.Customer_id)
	fmt.Println(o.Delivery_service)
	fmt.Println(o.Shardkey)
	fmt.Println(o.Sm_id)
	fmt.Println(o.Date_created)
	fmt.Println(o.Oof_shard)
}
