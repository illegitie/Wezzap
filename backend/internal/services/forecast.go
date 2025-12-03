package service

import (
	"wezap/internal/client"
	"wezap/internal/models"
)

type ForecastService struct {
	api *client.WeatherClient
}

func NewForecastService(api *client.WeatherClient) *ForecastService {
	return &ForecastService{
		api: api,
	}
}

func (s *ForecastService) GetCurrentWeather(place string) (models.CurrentWeatherForecast, error) {
	return s.api.GetCurrentWeather(place)
}

func (s *ForecastService) GetForecastEvery3Hours(city string) ([]models.Forecast, error) {
	return s.api.GetForecastEvery3Hours(city)
}

func (s *ForecastService) GetForecastPerDay(place string) ([]models.Forecast, error) {
	return s.api.GetForecastPerDay(place)
}
