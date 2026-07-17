package main

import (
	"church/routes"
	"log"
	"net/http"
)

func main() {
	routes.RegisterRoutes()
	log.Println("Server Activated....")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
	}
}
