package models

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type WeatherAPIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, StatusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(StatusCode, errorResponse{message})
}

type StatusResponse struct {
	Status string `json:"status"`
}

type OpenWeatherApiResponse struct {
	List []struct {
		DtTxt string `json:"dt_txt"`
		Main  struct {
			Temp    float64 `json:"temp"`
			TempMin float64 `json:"temp_min"`
			TempMax float64 `json:"temp_max"`
		} `json:"main"`
		Weather []WeatherCondition `json:"weather"`
	} `json:"list"`
}

type CurrentWeatherResponse struct {
	Name       string `json:"name"`
	Dt         int64  `json:"dt"`
	Visibility int    `json:"visibility"`
	Main       struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  float64 `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Weather []WeatherCondition `json:"weather"`
}
