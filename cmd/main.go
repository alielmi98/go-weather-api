package main

import (
	"github.com/alielmi98/go-weather-api/internal/cache"
	"github.com/alielmi98/go-weather-api/internal/config"
	"github.com/alielmi98/go-weather-api/internal/handlers"
	"github.com/alielmi98/go-weather-api/internal/services"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Initialize cache
	cache := cache.NewCache(cfg)

	// Initialize services
	weatherService := services.NewWeatherService(cfg, cache)

	// Initialize handlers
	weatherHandler := handlers.NewWeatherHandler(weatherService)

	// Set up Gin router
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/weather/:city", weatherHandler.GetWeatherByCity)

	// Start server
	r.Run(cfg.ServerPort)

}
