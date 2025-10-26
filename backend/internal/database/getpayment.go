package database

import (
	"log"
	"wb_labs_l0/backend/internal/model"
)

func GetPaymentByID(id int) []model.Payment {
	rows, err := DB.Query(`
	SELECT * FROM Payments
	WHERE id = $1
	`, id)
	if err != nil {
		log.Println(err)
	}
	payments := []model.Payment{}
	for rows.Next() {
		payment := model.Payment{}
		err = rows.Scan(
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
		payments = append(payments, payment)
	}
	return payments
}
func GetPaymentByTransaction(Transaction string) model.Payment {
	row := DB.QueryRow(`
	SELECT * FROM Payments
	WHERE name = $1
	`, Transaction)
	var empty string
	payment := model.Payment{}
	err := row.Scan(
		&empty,
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
func GetAllPayments() []model.Payment {
	rows, err := DB.Query(`
	SELECT * FROM Payments
	`)
	if err != nil {
		log.Println(err)
	}
	var empty string
	payments := []model.Payment{}
	for rows.Next() {
		payment := model.Payment{}
		err = rows.Scan(
			&empty,
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
		payments = append(payments, payment)
	}

	return payments
}
