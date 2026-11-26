package routes

import (
	"net/http"
	"station/controllers"
)

func MainRouter() *http.ServeMux {
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	router.Handle("GET /static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("GET /{$}", controllers.Home)
	router.HandleFunc("GET /ville", controllers.ListCommunes)
	

	return router
}