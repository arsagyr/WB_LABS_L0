package handlers

import (
	"log"
	"net/http"
)

func SearchOrder(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		id := r.FormValue(("id"))
		// orders := database.GetOrderByUID(id)

		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/items/"+id, 301)
	} else {
		http.ServeFile(w, r, "backend/templates/search.html")
	}

}
