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
	forecast, err := api.GetForecast(commune[0].Centre.Coordonnees[1], commune[0].Centre.Coordonnees[0])
	forecast.Nom = commune[0].Nom
	return forecast, err
}
