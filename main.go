package main

import (
	/* 	"fmt"
	   	"log"
	   	"net/http"
	   	"station/routes" */
	"fmt"
	"station/services"
)

func main() {
/* 	router := routes.MainRouter()
	
	fmt.Println("http://localhost:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Lancement du serveur a échoué: %v", err)
		return
	} */
	
	fmt.Println(services.GetCommunesByName("Salon"))
	
}