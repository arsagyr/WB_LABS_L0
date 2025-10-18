package model

// id SERIAL PRIMARY KEY,
// name VARCHAR(100) NOT NULL,
// phone VARCHAR(20),
// zip VARCHAR(20),
// city VARCHAR(100),
// address TEXT,
// region VARCHAR(100),
// email VARCHAR(100)

type Deliveries struct {
	Id      int    `json:"id" db:"id"`        //
	Name    string `json:"name" db:"name"`    //
	Phone   string `json:"phone" db:"name"`   //
	Zip     string `json:"zip" db:"name"`     //
	City    string `json:"city" db:"name"`    //
	Address string `json:"address" db:"name"` //
	Region  string `json:"region" db:"name"`  //
	Email   string `json:"email" db:"name"`   //
}
