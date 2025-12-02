package models

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type WeatherAPIResponse struct {
	Success bool       `json:"success"`
	Data    []Forecast `json:"data,omitempty"`
	Error   string     `json:"error"`
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
			TempMin float64 `json:"temp_min"`
			TempMax float64 `json:"temp_max"`
		} `json:"main"`
	} `json:"list"`
}
