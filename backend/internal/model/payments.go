package model

//	"payment": {
//		"transaction": "b563feb7b2b84b6test",
//		"request_id": "",
//		"currency": "USD",
//		"provider": "wbpay",
//		"amount": 1817,
//		"payment_dt": 1637907727,
//		"bank": "alpha",
//		"delivery_cost": 1500,
//		"goods_total": 317,
//		"custom_fee": 0
//	},
// type Payments struct {
// 	ID           int    `json:"id" db:"id"`
// 	Transaction  string `json:"transaction" db:"transaction"`
// 	RequestID    string `json:"request_id" db:"request_id"`
// 	Currency     string `json:"currency" db:"currency"`
// 	Provider     string `json:"provider" db:"provider"`
// 	Amount       string `json:"amount" db:"amount"`
// 	PaymentDt    string `json:"payment_dt" db:"payment_dt"`
// 	Bank         string `json:"bank" db:"bank"`
// 	DeliveryCost string `json:"delivery_cost" db:"delivery_cost"`
// 	GoodsTotal   string `json:"goods_total" db:"goods_total"`
// 	CustomFee    string `json:"custom_fee" db:"custom_fee"`
// }

type Payment struct {
	Transaction  string `json:"transaction" db:"transaction"`
	RequestID    string `json:"request_id" db:"request_id"`
	Currency     string `json:"currency" db:"currency"`
	Provider     string `json:"provider" db:"provider"`
	Amount       int    `json:"amount" db:"amount"`
	PaymentDt    int64  `json:"payment_dt" db:"payment_dt"`
	Bank         string `json:"bank" db:"bank"`
	DeliveryCost int    `json:"delivery_cost" db:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total" db:"goods_total"`
	CustomFee    int    `json:"custom_fee" db:"custom_fee"`
}
