package database

import (
	"log"
	"wb_labs_l0/backend/internal/model"
)

func GetDeliveryByID(id int) model.Delivery {
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
	var empty string
	delivery := model.Delivery{}
	err := row.Scan(
		&empty,
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

func GetAllDeliveries() []model.Delivery {
	rows, err := DB.Query(`
	SELECT * FROM Deliveries
	`)
	if err != nil {
		log.Println(err)
	}
	var empty string
	deliveries := []model.Delivery{}
	for rows.Next() {
		delivery := model.Delivery{}
		err = rows.Scan(
			&empty,
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
		deliveries = append(deliveries, delivery)
	}

	return deliveries
}
