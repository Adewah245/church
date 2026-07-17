package routes

import (
	"church/handlers"
	"church/middleware"
	"net/http"
)

func RegisterRoutes() {
	http.Handle(
		"/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")),
		),
	)
	http.HandleFunc("/", middleware.OnlyPath("/", middleware.OnlyGet(handlers.HomeHandler)))
	http.HandleFunc("/login", middleware.OnlyGet(handlers.LoginHandler))
	http.HandleFunc("/register", middleware.OnlyGet(handlers.RegisterHandler))

}
