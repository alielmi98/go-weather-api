package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/alielmi98/go-weather-api/internal/cache"
	"github.com/alielmi98/go-weather-api/internal/config"
)

type WeatherService interface {
	GetWeatherByCity(ctx context.Context, city string) (WeatherResponse, error)
}

type weatherService struct {
	apiKey string
	cache  cache.Cache
}

type WeatherResponse struct {
	Location    string    `json:"location"`
	Temperature float64   `json:"temperature"`
	Conditions  string    `json:"conditions"`
	Time        time.Time `json:"time"`
}

func NewWeatherService(cfg *config.Config, cache *cache.Cache) WeatherService {
	return &weatherService{
		apiKey: cfg.WeatherAPIKey,
		cache:  *cache,
	}
}

func (s *weatherService) GetWeatherByCity(ctx context.Context, city string) (WeatherResponse, error) {
	cacheKey := fmt.Sprintf("weather:%s", city)

	// Check cache
	cachedWeather, err := s.cache.Get(ctx, cacheKey)
	if err == nil && cachedWeather != "" {
		var resp WeatherResponse
		err = json.Unmarshal([]byte(cachedWeather), &resp)
		if err == nil {
			return resp, nil
		}
	}

	// Fetch weather from API
	resp, err := http.Get(fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&key=%s&contentType=json", city, s.apiKey))
	if err != nil {
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()

	var apiResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return WeatherResponse{}, err
	}

	// Extract weather data
	weatherData := apiResponse["days"].([]interface{})[0].(map[string]interface{})
	weatherResp := WeatherResponse{
		Location:    city,
		Temperature: weatherData["temp"].(float64),
		Conditions:  weatherData["conditions"].(string),
		Time:        time.Now(),
	}

	// Cache the response
	weatherJSON, _ := json.Marshal(weatherResp)
	err = s.cache.Set(ctx, cacheKey, string(weatherJSON), 12*time.Hour)
	if err != nil {
		return WeatherResponse{}, err
	}

	return weatherResp, nil
}
