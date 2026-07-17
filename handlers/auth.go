package handlers

import (
	"church/render"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	log.Println("User visited Login Portal")
	if err := render.Render(w, "login.html", nil); err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
}