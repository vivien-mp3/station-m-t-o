package models

type Forecast struct {
	Nom string
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone string `json:"timezone"`
	Timezone_Abbreviation string `json:"timezone_abbreviation"`
	Elevation float64 `json:"elevation"`
	Daily_Units struct{
		Time string `json:"time"`
		TempMax string `json:"temperature_2m_max"`
		TempMin string `json:"temperature_2m_min"`
		Weather_Code string `json:"weather_code"`
	} `json:"daily_units"`
	Daily struct{
		Time []string `json:"time"`
		TempMax []float64 `json:"temperature_2m_max"`
		TempMin []float64 `json:"temperature_2m_min"`
		Weather_Code []int `json:"weather_code"`
	} `json:"daily"`
} 