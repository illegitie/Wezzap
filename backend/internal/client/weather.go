package client

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"sort"
	"strings"
	"time"
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

func (c *WeatherClient) GetCurrentWeather(place string) (models.CurrentWeatherForecast, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", place, c.apiKey)
	resp, err := c.Client.Get(url)
	if err != nil {
		return models.CurrentWeatherForecast{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.CurrentWeatherForecast{}, fmt.Errorf("status code %d", resp.StatusCode)
	}
	var raw models.CurrentWeatherResponse
	if err = json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return models.CurrentWeatherForecast{}, err
	}
	result := models.CurrentWeatherForecast{
		Day:        time.Unix(raw.Dt, 0).Format("2006-01-02 15:04:05"),
		Temp:       raw.Main.Temp,
		Max:        raw.Main.TempMax,
		Min:        raw.Main.TempMin,
		Name:       raw.Name,
		FeelsLike:  raw.Main.FeelsLike,
		Pressure:   raw.Main.Pressure,
		Humidity:   raw.Main.Humidity,
		Visibility: raw.Visibility,
		WindSpeed:  raw.Wind.Speed,
		WindDeg:    raw.Wind.Deg,
		Clouds:     raw.Clouds.All,
		Weather:    raw.Weather,
	}
	return result, nil
}

func (c *WeatherClient) GetForecastEvery3Hours(city string) ([]models.ForecastEvery3Hours, error) {
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

	result := make([]models.ForecastEvery3Hours, 0, len(raw.List))
	for _, item := range raw.List {
		result = append(result, models.ForecastEvery3Hours{
			Time:    item.DtTxt,
			Temp:    item.Main.Temp,
			Weather: item.Weather,
		})
	}
	return result, nil
}

func (c *WeatherClient) GetForecastPerDay(place string) ([]models.Forecast, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?q=%s&appid=%s&units=metric", place, c.apiKey)

	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}
	var raw struct {
		List []struct {
			DtTxt string `json:"dt_txt"`
			Main  struct {
				TempMin float64 `json:"temp_min"`
				TempMax float64 `json:"temp_max"`
			} `json:"main"`
			Weather []models.WeatherCondition `json:"weather"`
		} `json:"list"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}
	dailyMap := make(map[string]models.Forecast)

	for _, item := range raw.List {
		day := strings.Split(item.DtTxt, " ")[0]
		if val, exists := dailyMap[day]; exists {
			val.Min = math.Min(val.Min, item.Main.TempMin)
			val.Max = math.Max(val.Max, item.Main.TempMax)
			dailyMap[day] = val
		} else {
			dailyMap[day] = models.Forecast{
				Day:     day,
				Min:     item.Main.TempMin,
				Max:     item.Main.TempMax,
				Weather: item.Weather,
			}
		}
	}

	dailyForecast := make([]models.Forecast, 0, len(dailyMap))
	for _, item := range dailyMap {
		dailyForecast = append(dailyForecast, item)
	}
	sort.Slice(dailyForecast, func(i, j int) bool {
		return dailyForecast[i].Day < dailyForecast[j].Day
	})

	return dailyForecast, nil
}
