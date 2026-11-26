package routes

import (
	"net/http"
	"station/controllers"
)

func MainRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /{$}", controllers.ListCommunes)

	return router
}