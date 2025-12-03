package handler

import (
	"net/http"
	"wezap/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetForecastEvery3Hours(c *gin.Context) {
	place := c.Query("place")
	if place == "" {
		models.NewErrorResponse(c, http.StatusBadRequest, "place is empty")
		return
	}

	forecast, err := h.services.GetForecastEvery3Hours(place)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, "error getting forecast")
		return
	}

	c.JSON(http.StatusOK, models.WeatherAPIResponse{
		Success: true,
		Data:    forecast,
	})
}

func (h *Handler) GetCurrentWeather(c *gin.Context) {
	place := c.Query("place")
	if place == "" {
		models.NewErrorResponse(c, http.StatusBadRequest, "place is empty")
		return
	}
	weather, err := h.services.GetCurrentWeather(place)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, "error getting current weather")
		return
	}

	c.JSON(http.StatusOK, models.WeatherAPIResponse{
		Success: true,
		Data:    weather,
	})

}

func (h *Handler) GetForecastPerDay(c *gin.Context) {
	place := c.Query("place")
	if place == "" {
		models.NewErrorResponse(c, http.StatusBadRequest, "place is empty")
		return
	}

	forecast, err := h.services.GetForecastPerDay(place)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, "error getting forecast")
		return
	}
	c.JSON(http.StatusOK, models.WeatherAPIResponse{
		Success: true,
		Data:    forecast,
	})
}
