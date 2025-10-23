package handlers

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func CreateOrder() {

// }

// // функция добавления данных
// func CreateHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {
// 		err := r.ParseForm()
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		familyname := r.FormValue("familyname")
// 		givenname := r.FormValue("givenname")
// 		nation := "%" + r.FormValue("nation") + "%"
// 		number := r.FormValue("number")
// 		honorar := r.FormValue("honorar")

// 		fmt.Println(familyname, givenname, nation, number, honorar)

// 		_, err = DB.Exec(`
// 			INSERT INTO Names (Family, Given)
// 			VALUES ($1, $2);
// 		`, familyname, givenname)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		row := DB.QueryRow(`
// 		SELECT id FROM Nations WHERE Name LIKE $1;
// 		`, nation)
// 		nationid := 0
// 		err = row.Scan(&nationid)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		_, err = DB.Exec(`
// 			INSERT INTO Actors (nameid, nationid, number, honorar)
//  			VALUES  (
// 			(SELECT COALESCE(MAX(Id), 0) FROM  Names),
// 			$1, $2, $3
// 			);
// 			`, nationid, number, honorar)

// 		if err != nil {
// 			log.Println(err)
// 		}
// 		http.Redirect(w, r, "/", 301)
// 	} else {
// 		http.ServeFile(w, r, "frontend/templates/create.html")
// 	}
// }
