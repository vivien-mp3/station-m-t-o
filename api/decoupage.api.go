package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var client = http.Client{
	Timeout: 3 * time.Second,
}

func get(endpoint string, dest any) {

	request, requestErr := http.NewRequest(http.MethodGet, "https://geo.api.gouv.fr/departements/01/communes", nil)
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
		
	}

}