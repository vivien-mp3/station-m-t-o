package main

import (
	"fmt"
	"log"
	"net/http"
	"station/routes"
	"station/templates"
)

func main() {
	templates.LoadTemplates()
	router := routes.MainRouter()
	
	fmt.Println("http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatalf("Lancement du serveur a échoué: %v", err)
		return
	}
}