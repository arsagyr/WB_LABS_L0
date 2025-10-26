package model

// id SERIAL PRIMARY KEY,
// order_uid VARCHAR(50),
// track_number VARCHAR(50) NOT NULL,
// entry VARCHAR(10) NOT NULL,
// deliveries_id INTEGER,
// payments_id INTEGER,
// items_id INTEGER,
// locale VARCHAR(10),
// internal_signature VARCHAR(255),
// customer_id VARCHAR(50),
// delivery_service VARCHAR(50),
// shardkey VARCHAR(10),
// sm_id INTEGER,
// date_created TIMESTAMP WITH TIME ZONE,
// oof_shard VARCHAR(10)

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
