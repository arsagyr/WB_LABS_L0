package model

import "fmt"

type Delivery struct {
	Name    string `json:"name" db:"name"`    //
	Phone   string `json:"phone" db:"name"`   //
	Zip     string `json:"zip" db:"name"`     //
	City    string `json:"city" db:"name"`    //
	Address string `json:"address" db:"name"` //
	Region  string `json:"region" db:"name"`  //
	Email   string `json:"email" db:"name"`   //
}

func (d Delivery) Print() {
	fmt.Println(d.Name)
	fmt.Println(d.Phone)
	fmt.Println(d.Zip)
	fmt.Println(d.City)
	fmt.Println(d.Address)
	fmt.Println(d.Region)
	fmt.Println(d.Email)
}
