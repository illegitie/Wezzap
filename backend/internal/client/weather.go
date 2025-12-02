package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wezap/internal/models"
)

type WeatherClient struct {
	apiKey string
	Client *http.Client
}

func NewWeatherClient(apiKey string) *WeatherClient {
	return &WeatherClient{
		apiKey: apiKey,
		Client: http.DefaultClient,
	}
}

func (c *WeatherClient) GetForecast(city string) ([]models.Forecast, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?q=%s&appid=%s&units=metric", city, c.apiKey)

	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	var raw models.OpenWeatherApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	result := make([]models.Forecast, 0, len(raw.List))
	for _, item := range raw.List {
		result = append(result, models.Forecast{
			Day: item.DtTxt,
			Max: item.Main.TempMax,
			Min: item.Main.TempMin,
		})
	}
	return result, nil
}
