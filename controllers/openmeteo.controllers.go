package controllers

import (
	"log"
	"net/http"
	"station/services"
	"station/templates"
)


func ShowMeteo(w http.ResponseWriter, r *http.Request) {
	forecast, err := services.GetForecast(r.PathValue("id"))

	if err != nil {
		log.Printf("Erreur Meteo: %v", err)
	}

	templates.Render(w, http.StatusOK, "meteo", forecast)
}