package handlers

import (
	"church/render"
	"log"
	"net/http"
)

func MembersHandlers(w http.ResponseWriter, r *http.Request) {
	log.Println("User Visited Members Page")
	if err := render.Render(w, "members.html", nil); err != nil {
		log.Println("Render Can't Load", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}