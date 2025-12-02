package handler

import (
	"wezap/internal/services"

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

	forecast := router.Group("/forecast")
	{
		forecast.GET("/", h.GetForecast)
	}

	return router
}
