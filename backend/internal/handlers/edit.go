package handlers

// import (
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"net/http"
// 	model "wb_labs_l0/backend/internal/model"

// 	"github.com/gorilla/mux"
// )

// func EditOrder() {

// }

// func EditPage(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	row := DB.QueryRow(`
// 		SELECT Actors.id, Names.Family, Names.Given, Nations.Name, Number, Honorar FROM Actors
// 		JOIN Names ON Actors.Nameid=Names.id
// 		JOIN Nations ON Actors.Nationid=Nations.id
// 		WHERE Actors.id = $1
// 	`, id)
// 	a := model.Actor{}
// 	err := row.Scan(&a.Id, &a.Familyname, &a.Givenname, &a.Nation, &a.Number, &a.Honorar)
// 	if err != nil {
// 		log.Println(err)
// 		http.Error(w, http.StatusText(404), http.StatusNotFound)
// 	} else {
// 		tmpl, _ := template.ParseFiles("backend/templates/edit.html")
// 		tmpl.Execute(w, a)
// 	}
// }

// // получаем измененные данные и сохраняем их в БД
// func EditHandler(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	id := r.FormValue(("id"))
// 	familyname := r.FormValue("familyname")
// 	givenname := r.FormValue("givenname")
// 	nation := r.FormValue("nation")
// 	number := r.FormValue("number")
// 	honorar := r.FormValue("honorar")

// 	fmt.Println(id, familyname, givenname, nation, number, honorar)

// 	_, err = DB.Exec(`
// 	UPDATE Names SET Family= $1, Given=$2 WHERE id=(
// 	SELECT nameid FROM Actors WHERE id=$3
// 	)
// 	`,
// 		familyname, givenname, id)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	_, err = DB.Exec(`
// 	UPDATE Actors SET Nationid = (
// 		SELECT id FROM Nations
// 		WHERE Name LIKE $1
// 	),
// 	Number=$2,
// 	Honorar=$3
// 	WHERE id = $4
// 	`, nation, number, honorar, id)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	http.Redirect(w, r, "/", 301)
// }
