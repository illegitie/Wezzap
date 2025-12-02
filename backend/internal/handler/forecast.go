package handler

import (
	"net/http"
	"wezap/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetForecast(c *gin.Context) {
	place := c.Query("place")
	if place == "" {
		models.NewErrorResponse(c, http.StatusBadRequest, "place is empty")
		return
	}

	forecast, err := h.services.GetForecast(place)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, "error getting forecast")
		return
	}

	c.JSON(http.StatusOK, models.WeatherAPIResponse{
		Success: true,
		Data:    forecast,
	})
}
