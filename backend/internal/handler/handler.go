package handler

import (
	"wezap/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(s *service.Services) *Handler {
	return &Handler{
		services: s,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(cors.Default())
	forecast := router.Group("/forecast")
	{
		forecast.GET("/", h.GetForecastEvery3Hours)
		forecast.GET("/current", h.GetCurrentWeather)
		forecast.GET("/per-day", h.GetForecastPerDay)
	}

	return router
}
