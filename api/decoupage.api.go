package api

import (
	"encoding/json"
	"log"
	"net/http"
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
	}

	response, responseErr := client.Do(request)
	if responseErr != nil {
		log.Printf("errooooor")
	}

	defer response.Body.Close()

	err := json.NewDecoder(response.Body).Decode(dest)
	if err != nil {
		log.Printf("System32/delete")
	}

	return nil
}


func GetCommune() {

}