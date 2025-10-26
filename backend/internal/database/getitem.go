package database

import (
	"log"
	"wb_labs_l0/backend/internal/model"
)

func GetItemsByOrderID(id int) []model.Item {
	rows, err := DB.Query(`
	SELECT * FROM Items
	WHERE order_id = $1
	`, id)
	if err != nil {
		log.Println(err)
	}
	items := []model.Item{}
	for rows.Next() {
		item := model.Item{}
		err := rows.Scan(
			&id,
			&item.Chrt_id,
			&item.Track_number,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.Total_price,
			&item.Nm_id,
			&item.Brand,
			&item.Status,
			&id,
		)
		if err != nil {
			log.Println(err)
		}
		items = append(items, item)
	}

	return items
}

func GetItemsByOrderStringID(id string) []model.Item {
	rows, err := DB.Query(`
	SELECT * FROM Items
	WHERE order_id = $1
	`, id)
	if err != nil {
		log.Println(err)
	}
	items := []model.Item{}
	for rows.Next() {
		item := model.Item{}
		err := rows.Scan(
			&id,
			&item.Chrt_id,
			&item.Track_number,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.Total_price,
			&item.Nm_id,
			&item.Brand,
			&item.Status,
			&id,
		)
		if err != nil {
			log.Println(err)
		}
		items = append(items, item)
	}

	return items
}

func GetItemByChrt_id(id int) model.Item {
	row := DB.QueryRow(`
	SELECT * FROM Items
	WHERE chrt_id = $1
	`, id)
	item := model.Item{}
	err := row.Scan(
		&id,
		&item.Chrt_id,
		&item.Track_number,
		&item.Price,
		&item.Rid,
		&item.Name,
		&item.Sale,
		&item.Size,
		&item.Total_price,
		&item.Nm_id,
		&item.Brand,
		&item.Status,
		&id,
	)
	if err != nil {
		log.Println(err)
	}

	return item
}

func GetItemByID(id string) model.Item {
	row := DB.QueryRow(`
	SELECT * FROM Items
	WHERE id = $1
	`, id)
	item := model.Item{}
	err := row.Scan(
		&id,
		&item.Chrt_id,
		&item.Track_number,
		&item.Price,
		&item.Rid,
		&item.Name,
		&item.Sale,
		&item.Size,
		&item.Total_price,
		&item.Nm_id,
		&item.Brand,
		&item.Status,
		&id,
	)
	if err != nil {
		log.Println(err)
	}

	return item
}

func GetItems() []model.Item {
	rows, err := DB.Query(`
	SELECT * FROM Items
	`)
	if err != nil {
		log.Println(err)
	}
	var id string
	items := []model.Item{}
	for rows.Next() {
		item := model.Item{}
		err := rows.Scan(
			&id,
			&item.Chrt_id,
			&item.Track_number,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.Total_price,
			&item.Nm_id,
			&item.Brand,
			&item.Status,
			&id,
		)
		if err != nil {
			log.Println(err)
		}
		items = append(items, item)
	}

	return items
}
