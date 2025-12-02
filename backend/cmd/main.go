package main

import (
	"log"
	"wezap/internal/client"
	"wezap/internal/handler"
	backend "wezap/internal/server"
	"wezap/internal/services"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	apiKey := viper.GetString("client.apiKey")
	port := viper.GetString("port")
	srv := new(backend.Server)
	api := client.NewWeatherClient(apiKey)
	services := service.NewServices(api)
	handlers := handler.NewHandler(services)
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
