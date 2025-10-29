package database

import (
	"log"
	"wb_labs_l0/backend/internal/model"
)

func InsertItem(item model.Item) {

	_, err := DB.Exec(`
	INSERT INTO items (
    chrt_id, track_number, price, 
	rid, name,  sale, size, 
	total_price, nm_id, brand, 
	status, order_id) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11,
	(SELECT COALESCE(MAX(Id), 0) FROM  Orders))`,
		item.Chrt_id,
		item.Track_number,
		item.Price,
		item.Rid,
		item.Name,
		item.Sale,
		item.Size,
		item.Total_price,
		item.Nm_id,
		item.Brand,
		item.Status,
	)
	if err != nil {
		log.Println("Ошибка ввода товара: " + err.Error())
		_, err = DB.Exec(`
		DELETE FROM items
		WHERE order_id = (SELECT COALESCE(MAX(Id), 0) FROM  Orders);
		DELETE FROM  orders
		WHERE id = (SELECT COALESCE(MAX(Id), 0) FROM  Orders);
		DELETE FROM  payments
		WHERE payments.id = (SELECT COALESCE(MAX(Id), 0) FROM  payments);
		DELETE FROM  deliveries
		WHERE deliveries.id = (SELECT COALESCE(MAX(Id), 0) FROM  deliveries);
		`)
		if err != nil {
			log.Println("Ошибка удаления данных заказа: " + err.Error())
		}
	}
}

func InsertDelivery(delivery model.Delivery) error {

	_, err := DB.Exec(`
	INSERT INTO deliveries (
    name, phone, 
    zip, city, address, 
    region, email
	) VALUES (
	 $1, $2, $3, $4, $5, $6, $7)
	`,
		delivery.Name,
		delivery.Phone,
		delivery.Zip,
		delivery.City,
		delivery.Address,
		delivery.Region,
		delivery.Email,
	)
	if err != nil {
		log.Println("Ошибка ввода поставщика: " + err.Error())
	}
	return err
}

func InsertPayment(payment model.Payment) error {

	_, err := DB.Exec(`
	INSERT INTO payments (
	transaction, request_id, 
	currency, provider, amount, 
	payment_dt, bank, delivery_cost, 
	goods_total,  custom_fee
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		payment.Transaction,
		payment.RequestID,
		payment.Currency,
		payment.Provider,
		payment.Amount,
		payment.PaymentDt,
		payment.Bank,
		payment.DeliveryCost,
		payment.GoodsTotal,
		payment.CustomFee,
	)
	if err != nil {
		log.Println("Ошибка ввода оплаты: " + err.Error())
	}
	return err
}

func InsertOrder(orderdb model.Order) error {
	err := InsertDelivery(orderdb.Delivery)
	if err != nil {
		goto end
	}
	err = InsertPayment(orderdb.Payment)
	if err != nil {
		goto end
	}
	_, err = DB.Exec(`
        INSERT INTO orders (
		order_uid, track_number, entry, 
		locale, internal_signature, customer_id, 
		delivery_service, shardkey, sm_id, 
		date_created, oof_shard, 
		delivery_id, payment_id 
		) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11,
		(SELECT COALESCE(MAX(Id), 0) FROM  Deliveries),
		(SELECT COALESCE(MAX(Id), 0) FROM  Payments))`,
		orderdb.Order_uid,
		orderdb.Track_number,
		orderdb.Entry,
		orderdb.Locale,
		orderdb.Internal_signature,
		orderdb.Customer_id,
		orderdb.Delivery_service,
		orderdb.Shardkey,
		orderdb.Sm_id,
		orderdb.Date_created,
		orderdb.Oof_shard,
	)

	if err != nil {
		log.Println("Ошибка ввода заказа: " + err.Error())
		_, err = DB.Exec(`
		DELETE FROM Deliveries
		WHERE id=(SELECT COALESCE(MAX(Id), 0) FROM  Deliveries)
		`)
		if err != nil {
			log.Println("Ошибка удаления данных поставщика: " + err.Error())
		}
		_, err = DB.Exec(`
		DELETE FROM Payments
		WHERE id=(SELECT COALESCE(MAX(Id), 0) FROM  Payments)
		`)
		if err != nil {
			log.Println("Ошибка удаления данных оплаты: " + err.Error())
		}
		goto end
	} else {
		for i := 0; i < len(orderdb.Items); i++ {
			InsertItem(orderdb.Items[i])
			goto end
		}
	}
end:
	return err
}
