package handlers

import (
	"html/template"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func CacheFileNameMaker() string {
	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	d := time.Now().Sub(start)
	s := d.String()
	s = "cache/" + s
	return s
}

func Mainpage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("frontend/templates/mainpage.html")
	tmpl.Execute(w, r)
}
