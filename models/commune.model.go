package models

type Commune struct {
	Code string `json:"code"`
	Nom string `json:"nom"`
	CodePostaux []string `json:"codePostaux"`
	Population int `json:"population"`
	Centre struct{
		Type string `json:"type"`
		Coordonnees []float64 `json:"coordinates"`
	} `json:"centre,omitempty"`
	Score float64 `json:"_score"`
} 