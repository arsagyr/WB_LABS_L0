package database

import (
	"log"
	"wb_labs_l0/backend/internal/model"
)

func GetDeliveryByID(id string) model.Delivery {
	row := DB.QueryRow(`
	SELECT * FROM Deliveries
	WHERE id = $1
	`, id)
	delivery := model.Delivery{}
	err := row.Scan(
		&id,
		&delivery.Name,
		&delivery.Phone,
		&delivery.Zip,
		&delivery.City,
		&delivery.Address,
		&delivery.Region,
		&delivery.Email,
	)
	if err != nil {
		log.Println(err)
	}

	return delivery
}

func GetDeliveryByName(name string) model.Delivery {
	row := DB.QueryRow(`
	SELECT * FROM Deliveries
	WHERE name = $1
	`, name)
	delivery := model.Delivery{}
	err := row.Scan(
		&delivery.Name,
		&delivery.Phone,
		&delivery.Zip,
		&delivery.City,
		&delivery.Address,
		&delivery.Region,
		&delivery.Email,
	)
	if err != nil {
		log.Println(err)
	}

	return delivery
}
