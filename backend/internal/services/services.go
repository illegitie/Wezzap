package service

import (
	"wezap/internal/client"
	"wezap/internal/models"
)

type Services struct {
	Forecast
}

type Forecast interface {
	GetForecast(place string) ([]models.Forecast, error)
}

func NewServices(api *client.WeatherClient) *Services {
	return &Services{
		NewForecastService(api),
	}
}
