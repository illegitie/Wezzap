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

func (s *ForecastService) GetForecast(city string) ([]models.Forecast, error) {
	return s.api.GetForecast(city)
}
