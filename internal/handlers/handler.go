package handler

import (
	"github.com/gin-gonic/gin"
	weatherapi "github.com/naumovrus/weather-api/internal/pkg/weather_api"
	"github.com/naumovrus/weather-api/internal/services"
)

type Handler struct {
	services   *services.Service
	weatherapi *weatherapi.WeatherAPI
}

func NewHandler(services *services.Service, weatherapi *weatherapi.WeatherAPI) *Handler {
	return &Handler{services: services, weatherapi: weatherapi}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		cityies := api.Group("/city")
		{
			cityies.POST("/", h.addCity)
			cityies.GET("/all", h.getAllCityies)
			cityies.GET("/", h.getUsersCity)
			// cityies.GET("/:name", h.getCityByName)

		}
		weather := api.Group("/weather")
		{
			weather.POST("/", h.GetWeather)
		}

	}
	return router
}
