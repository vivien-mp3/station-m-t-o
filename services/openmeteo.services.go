package services

import (
	"station/api"
	"station/models"
)

func GetForecast(code string) (models.Forecast, error) {
	commune, err := GetCommuneByCode(code)

	return api.GetForecast(6.4963, 48.788)
}
