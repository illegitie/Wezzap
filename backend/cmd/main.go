package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wezap/internal/client"
	"wezap/internal/handler"
	backend "wezap/internal/server"
	"wezap/internal/services"

	"github.com/sirupsen/logrus"
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

	srv := new(backend.Server)
	api := client.NewWeatherClient(apiKey)
	services := service.NewServices(api)
	handlers := handler.NewHandler(services)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("errorResponse occured while running http server: %s", err.Error())
		}
	}()

	logrus.Infof("starting http server on port %s", viper.GetString("port"))
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logrus.Fatalf("error occurred while shutting down server: %v", err)
	}

	logrus.Info("server exited properly")
}
