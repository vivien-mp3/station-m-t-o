package controllers

import (
	"net/http"
	"station/services"
	"station/templates"
)

func ListCommunes(w http.ResponseWriter, r *http.Request) {
	communes, _ := services.GetCommunes()

	templates.Render(w, http.StatusOK, "home", communes)
}