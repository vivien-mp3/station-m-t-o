package api

import (
	"encoding/json"

	"log"
	"net/http"
	"station/models"
	"time"
)

const (
	baseURL = "https://geo.api.gouv.fr/departements/01/communes"
)

var client = http.Client{
	Timeout: 3 * time.Second,
}

func get(endpoint string, dest any) error {

	request, requestErr := http.NewRequest(http.MethodGet, baseURL + endpoint, nil)
	if requestErr != nil {
		log.Printf("errorbjvjhh")
		return requestErr
	}

	response, responseErr := client.Do(request)
	if responseErr != nil {
		log.Printf("errooooor")
		return responseErr
	}

	defer response.Body.Close()

	err := json.NewDecoder(response.Body).Decode(dest)
	if err != nil {
		log.Printf("System32/delete %w", err)
		return err
	}

	return nil
}


func GetCommunes() ([]models.Commune, error) {
	var communes []models.Commune
	err := get("?nom=Pussy&fields=code,nom,codesPostaux,population,centre&boost=population&limit=10", &communes)
	return communes, err
}