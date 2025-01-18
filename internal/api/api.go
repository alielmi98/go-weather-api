package api

import (
	_ "github.com/alielmi98/go-weather-api/docs"
	"github.com/alielmi98/go-weather-api/internal/cache"
	"github.com/alielmi98/go-weather-api/internal/config"
	"github.com/alielmi98/go-weather-api/internal/handlers"
	middlewares "github.com/alielmi98/go-weather-api/internal/middleware"
	"github.com/alielmi98/go-weather-api/internal/services"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config, cache *cache.Cache) {

	// Initialize services
	weatherService := services.NewWeatherService(cfg, cache)

	// Initialize handlers
	weatherHandler := handlers.NewWeatherHandler(weatherService)
	// Set up Gin router
	r := gin.Default()
	r.Use(middlewares.LimitByRequest())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/weather/:city", weatherHandler.GetWeatherByCity)
	// Start server
	r.Run(cfg.ServerPort)
}
