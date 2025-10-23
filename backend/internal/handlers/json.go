package handlers

// func PostOrderJSON(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	item := model.Item{
// 		Id:           r.FormValue(""),
// 		Chrt_id:      r.FormValue(""),
// 		Track_number: r.FormValue(""),
// 		Price:        r.FormValue(""),
// 		Rid:          r.FormValue(""),
// 		Name:         r.FormValue(""),
// 		Sale:         r.FormValue(""),
// 		Size:         r.FormValue(""),
// 		Total_price:  r.FormValue(""),
// 		Nm_id:        r.FormValue(""),
// 		Brand:        r.FormValue(""),
// 		Status:       r.FormValue(""),
// 	}

// 	delivery := model.Delivery{}

// 	order := model.Order{
// 		Id:                 r.FormValue(""),
// 		Order_uid:          r.FormValue(""),
// 		Track_number:       r.FormValue(""),
// 		Entry:              r.FormValue(""),
// 		Deliveries:         delivery,
// 		Payments:           payment,
// 		Items:              items,
// 		Locale:             r.FormValue(""),
// 		Internal_signature: r.FormValue(""),
// 		Customer_id:        r.FormValue(""),
// 		Delivery_service:   r.FormValue(""),
// 		Shardkey:           r.FormValue(""),
// 		Sm_id:              r.FormValue(""),
// 		Date_created:       r.FormValue(""),
// 		Oof_shard:          r.FormValue(""),
// 	}

// 	fmt.Println(
// 		order.Id,
// 		order.Order_uid,
// 		order.Track_number,
// 		order.Entry,
// 		order.Deliveries_id,
// 		order.Payments_id,
// 		order.Items_id,
// 		order.Locale,
// 		order.Internal_signature,
// 		order.Customer_id,
// 		order.Delivery_service,
// 		order.Shardkey,
// 		order.Sm_id,
// 		order.Date_created,
// 		order.Oof_shard,
// 	)

// 	data, err := json.MarshalIndent(order, "", "  ")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	err = os.WriteFile("frontend/json/order.json", data, 0644)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func LoadOrderJSON() {

// 	file, err := os.ReadFile("frontend/json/order.json")
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	var order model.Order
// 	err = json.Unmarshal(file, &order)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	_, err = DB.Exec(`
// 			INSERT INTO Orders (Family, Given)
// 			VALUES ($1, $2);
// 		`, order.Familyname, order.Givenname)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	row := DB.QueryRow(`
// 		SELECT id FROM Nations WHERE Name LIKE $1;
// 		`, order.Nation)
// 	nationid := 0
// 	err = row.Scan(&nationid)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	_, err = DB.Exec(`
// 			INSERT INTO Actors (nameid, nationid, number, honorar)
//  			VALUES  (
// 			(SELECT COALESCE(MAX(Id), 0) FROM  Names),
// 			$1, $2, $3
// 			);
// 			`, order.Nation, order.Number, order.Honorar)

// 	if err != nil {
// 		log.Println(err)
// 	}

// }

// func MakeOrderJSON() {
// 	item := model.Item{
// 		Id:           ,
// 		Chrt_id:      ,
// 		Track_number: ,
// 		Price:        r.FormValue(""),
// 		Rid:          r.FormValue(""),
// 		Name:         r.FormValue(""),
// 		Sale:         r.FormValue(""),
// 		Size:         r.FormValue(""),
// 		Total_price:  r.FormValue(""),
// 		Nm_id:        r.FormValue(""),
// 		Brand:        r.FormValue(""),
// 		Status:       r.FormValue(""),
// 	}

// 	delivery := model.Delivery{}

// 	order := model.Order{
// 		Id:                 r.FormValue(""),
// 		Order_uid:          r.FormValue(""),
// 		Track_number:       r.FormValue(""),
// 		Entry:              r.FormValue(""),
// 		Deliveries:         delivery,
// 		Payments:           payment,
// 		Items:              items,
// 		Locale:             r.FormValue(""),
// 		Internal_signature: r.FormValue(""),
// 		Customer_id:        r.FormValue(""),
// 		Delivery_service:   r.FormValue(""),
// 		Shardkey:           r.FormValue(""),
// 		Sm_id:              r.FormValue(""),
// 		Date_created:       r.FormValue(""),
// 		Oof_shard:          r.FormValue(""),
// 	}

// 	data, err := json.MarshalIndent(order, "", "  ")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	err = os.WriteFile("frontend/json/order.json", data, 0644)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }
