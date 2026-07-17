package handlers

import (
	"church/render"
	"log"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("User visited Register port")
	if err := render.Render(w, "register.html", nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}