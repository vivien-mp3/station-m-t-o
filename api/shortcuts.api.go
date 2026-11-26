package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	dcpURL = "https://geo.api.gouv.fr/communes"
	omURL = "https://api.open-meteo.com/v1/forecast"
)

var client = http.Client{
	Timeout: 5 * time.Second,
}

func get(URL, endpoint string, dest any) error {

	request, requestErr := http.NewRequest(http.MethodGet, URL + endpoint, nil)
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
		log.Printf("System32/delete %v", err)
		return err
	}

	return nil
}