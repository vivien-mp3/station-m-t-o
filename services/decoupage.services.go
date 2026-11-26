package services

import (
	"station/api"
	"station/models"
)

func GetCommunes() ([]models.Commune, error) {
	return api.GetCommunes()
}

func GetCommunesByName(name string) ([]models.Commune, error) {
	return api.GetCommunesByName(name)
}

func GetCommuneByCode(code string) (models.Commune, error) {
	return api.GetCommuneByCode(code)
}