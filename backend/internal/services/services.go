package service

import (
	"wezap/internal/client"
	"wezap/internal/models"
)

type Services struct {
	Forecast
}

type Forecast interface {
	GetForecastEvery3Hours(place string) ([]models.ForecastEvery3Hours, error)
	GetCurrentWeather(place string) (models.CurrentWeatherForecast, error)
	GetForecastPerDay(place string) ([]models.Forecast, error)
}

func NewServices(api *client.WeatherClient) *Services {
	return &Services{
		NewForecastService(api),
	}
}
