package api

import (
	"fmt"
	"station/models"
)

func GetCommunes() ([]models.Commune, error) {
	var communes []models.Commune
	err := get(dcpURL, "?fields=code,nom,codesPostaux,population,centre&boost=population", &communes)
	return communes, err
}

func GetCommunesByName(name string) ([]models.Commune, error) {
	var communes []models.Commune
	err := get(dcpURL, fmt.Sprintf("?nom=%s&fields=code,nom,codes,codesPostaux,population,centre&boost=population&limit=10", name), &communes)
	return communes, err
}

func GetCommuneByCode(code string) ([]models.Commune, error) {
	var commune []models.Commune
	err := get(dcpURL, fmt.Sprintf("?code=%s&fields=code,nom,codes,codesPostaux,population,centre&boost=population", code), &commune)
	return commune, err
}
