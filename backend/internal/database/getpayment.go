package database

import (
	"log"
	"wb_labs_l0/backend/internal/model"
)

func GetPaymentByID(id string) model.Payment {
	row := DB.QueryRow(`
	SELECT * FROM Payments
	WHERE id = $1
	`, id)
	payment := model.Payment{}
	err := row.Scan(
		&id,
		&payment.Transaction,
		&payment.RequestID,
		&payment.Currency,
		&payment.Provider,
		&payment.Amount,
		&payment.PaymentDt,
		&payment.Bank,
		&payment.DeliveryCost,
		&payment.GoodsTotal,
		&payment.CustomFee,
	)
	if err != nil {
		log.Println(err)
	}

	return payment
}
func GetPaymentByTransaction(Transaction string) model.Payment {
	row := DB.QueryRow(`
	SELECT * FROM Payments
	WHERE name = $1
	`, Transaction)
	payment := model.Payment{}
	err := row.Scan(
		&payment.Transaction,
		&payment.RequestID,
		&payment.Currency,
		&payment.Provider,
		&payment.Amount,
		&payment.PaymentDt,
		&payment.Bank,
		&payment.DeliveryCost,
		&payment.GoodsTotal,
		&payment.CustomFee,
	)
	if err != nil {
		log.Println(err)
	}

	return payment
}
