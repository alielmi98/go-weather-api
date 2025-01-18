package main

import (
	"github.com/alielmi98/go-weather-api/internal/api"
	"github.com/alielmi98/go-weather-api/internal/cache"
	"github.com/alielmi98/go-weather-api/internal/config"

	_ "github.com/alielmi98/go-weather-api/docs"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Initialize cache
	cache := cache.NewCache(cfg)
	api.InitServer(cfg, &cache)

}
