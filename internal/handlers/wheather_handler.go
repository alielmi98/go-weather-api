package handlers

import (
	"net/http"

	"github.com/alielmi98/go-weather-api/internal/services"

	"github.com/gin-gonic/gin"
)

type WeatherHandler struct {
	weatherService services.WeatherService
}

func NewWeatherHandler(weatherService services.WeatherService) *WeatherHandler {
	return &WeatherHandler{
		weatherService: weatherService,
	}
}

// GetWeatherByCity godoc
// @Summary Get weather information for a city
// @Description Get the current weather information for a given city.
// @Produce  json
// @Param city path string true "City name"
// @Success 200 {object} services.WeatherResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /weather/{city} [get]
func (h *WeatherHandler) GetWeatherByCity(c *gin.Context) {
	city := c.Param("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "missing city parameter",
		})
		return
	}

	weather, err := h.weatherService.GetWeatherByCity(c.Request.Context(), city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, weather)

}
