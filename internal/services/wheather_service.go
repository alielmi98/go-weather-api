package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/alielmi98/go-weather-api/internal/cache"
	"github.com/alielmi98/go-weather-api/internal/config"
	"github.com/alielmi98/go-weather-api/internal/dto"
)

type WeatherService interface {
	GetWeatherByCity(ctx context.Context, city dto.CityRequest) (dto.WeatherResponse, error)
}

type weatherService struct {
	apiKey string
	cache  cache.Cache
}

func NewWeatherService(cfg *config.Config, cache *cache.Cache) WeatherService {
	return &weatherService{
		apiKey: cfg.WeatherAPIKey,
		cache:  *cache,
	}
}

func (s *weatherService) GetWeatherByCity(ctx context.Context, city dto.CityRequest) (dto.WeatherResponse, error) {
	cacheKey := fmt.Sprintf("weather:%s", city.City)

	// Check cache
	cachedWeather, err := s.cache.Get(ctx, cacheKey)
	if err == nil && cachedWeather != "" {
		var resp dto.WeatherResponse
		err = json.Unmarshal([]byte(cachedWeather), &resp)
		if err == nil {
			return resp, nil
		}
	}

	// Fetch weather from API
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&key=%s&contentType=json", city.City, s.apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return dto.WeatherResponse{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.WeatherResponse{}, fmt.Errorf("failed to read response body: %v", err)
	}

	// Check if the status code is not 200 OK
	if resp.StatusCode != http.StatusOK {
		return dto.WeatherResponse{}, fmt.Errorf("external API request failed with status %d and error: %s", resp.StatusCode, string(respBody))
	}

	var apiResponse map[string]interface{}
	err = json.Unmarshal(respBody, &apiResponse)
	if err != nil {
		return dto.WeatherResponse{}, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	// Extract weather data
	weatherData := apiResponse["days"].([]interface{})[0].(map[string]interface{})
	location := apiResponse["resolvedAddress"].(string)
	weatherResp := dto.WeatherResponse{
		Location:    location,
		Temperature: weatherData["temp"].(float64),
		Windspeed:   weatherData["windspeed"].(float64),
		Conditions:  weatherData["conditions"].(string),
		Time:        time.Now(),
	}

	// Cache the response
	weatherJSON, _ := json.Marshal(weatherResp)
	err = s.cache.Set(ctx, cacheKey, string(weatherJSON), 12*time.Hour)
	if err != nil {
		return dto.WeatherResponse{}, err
	}

	return weatherResp, nil
}
