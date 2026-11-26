package controllers

import (
	"log"
	"net/http"
	"station/models"
	"station/services"
	"station/templates"
	"strings"
)

func Home(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, http.StatusOK, "home", nil)
}

func ListCommunes(w http.ResponseWriter, r *http.Request) {
	var communes []models.Commune 
	var err error

	query := r.URL.Query().Get("nom")
	log.Println(query)

	query = strings.ToLower(strings.ReplaceAll(query, " ", ""))

	if len(query) >= 2 {
		communes, err = services.GetCommunesByName(query)
	} else {
		communes, err = services.GetCommunes()
	}

	if err != nil {
		log.Printf("bru")
	}

	templates.Render(w, http.StatusOK, "ville", communes)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	templates.RenderError(w, http.StatusNotFound, "C'est page n'existe pas")
}