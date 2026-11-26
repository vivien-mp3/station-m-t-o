package services

import (
	"station/api"
	"station/models"
)

func GetCommunes() ([]models.Commune, error) {
	return api.GetCommunes()
}