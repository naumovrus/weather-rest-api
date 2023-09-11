package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/naumovrus/weather-api/internal/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
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

			// lists.PUT("/:id", h.updateList)
			// lists.DELETE("/:id", h.deleteList)

		}
		// weather := api.Group("/weather")
		// {
		// 	weather.GET("/", h.GetWeather)
		// }

	}
	return router
}
