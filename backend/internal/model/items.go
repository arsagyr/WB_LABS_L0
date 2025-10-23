package model

import "fmt"

// id SERIAL PRIMARY KEY,
// chrt_id INTEGER NOT NULL,
// track_number VARCHAR(50),
// price INTEGER,
// rid VARCHAR(100),
// name VARCHAR(255),
// sale INTEGER,
// size VARCHAR(10),
// total_price INTEGER,
// nm_id INTEGER,
// brand VARCHAR(100),
// status INTEGER

// type Item struct {
// 	Id           int    `json:"id" db:"id"` // ID
// 	Chrt_id      string `json:"chrt_id" db:"chrt_id"`
// 	Track_number string `json:"track_number" db:"track_number"`
// 	Price        string `json:"price" db:"price"`
// 	Rid          string `json:"rid" db:"rid"`
// 	Name         string `json:"name" db:"name"`
// 	Sale         string `json:"sale" db:"sale"`
// 	Size         string `json:"size" db:"size"`
// 	Total_price  string `json:"total_price" db:"total_price"`
// 	Nm_id        string `json:"nm_id" db:"nm_id"`
// 	Brand        string `json:"brand" db:"brand"`
// 	Status       string `json:"status" db:"status"`
// }

type Item struct {
	Chrt_id      int32  `json:"chrt_id" db:"chrt_id"`
	Track_number string `json:"track_number" db:"track_number"`
	Price        int    `json:"price" db:"price"`
	Rid          string `json:"rid" db:"rid"`
	Name         string `json:"name" db:"name"`
	Sale         int    `json:"sale" db:"sale"`
	Size         string `json:"size" db:"size"`
	Total_price  int    `json:"total_price" db:"total_price"`
	Nm_id        int32  `json:"nm_id" db:"nm_id"`
	Brand        string `json:"brand" db:"brand"`
	Status       int    `json:"status" db:"status"`
}

func (i Item) Print() {
	fmt.Println(i.Chrt_id)
	fmt.Println(i.Track_number)
	fmt.Println(i.Price)
	fmt.Println(i.Rid)
	fmt.Println(i.Name)
	fmt.Println(i.Sale)
	fmt.Println(i.Size)
	fmt.Println(i.Total_price)
	fmt.Println(i.Nm_id)
	fmt.Println(i.Brand)
	fmt.Println(i.Status)
}
