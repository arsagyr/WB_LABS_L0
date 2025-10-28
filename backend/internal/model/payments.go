package model

import "fmt"

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

func (p Payment) Print() {
	fmt.Println(p.Transaction)
	fmt.Println(p.RequestID)
	fmt.Println(p.Currency)
	fmt.Println(p.Provider)
	fmt.Println(p.Amount)
	fmt.Println(p.PaymentDt)
	fmt.Println(p.Bank)
	fmt.Println(p.DeliveryCost)
	fmt.Println(p.GoodsTotal)
	fmt.Println(p.CustomFee)
}
