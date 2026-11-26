package api

import (
	"fmt"
	"station/models"
)


func GetForecast(latitude, longitude float64) (models.Forecast, error) {
		var forecast models.Forecast
		err := get(omURL, fmt.Sprintf("?latitude=%f&longitude=%f&daily=temperature_2m_max,temperature_2m_min,weather_code&timezone=auto&forecast_days=7", latitude, longitude), &forecast)
		return forecast, err
}

