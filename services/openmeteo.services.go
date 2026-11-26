package services

import (
	"log"
	"station/api"
	"station/models"
)

func GetForecast(code string) (models.Forecast, error) {
	commune, err := GetCommuneByCode(code)
	if err != nil {
		log.Printf("Une erreur est parvenue. %v", err)
		return models.Forecast{}, err
	}
	return api.GetForecast(commune.Centre.Coordonnees[1], commune.Centre.Coordonnees[0])
}
