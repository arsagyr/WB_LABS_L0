package handlers

import (
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func Mainpage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("frontend/templates/mainpage.html")
	tmpl.Execute(w, r)
}
