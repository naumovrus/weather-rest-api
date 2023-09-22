package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	ent "github.com/naumovrus/weather-api/internal/entities"
)

func (h *Handler) GetWeather(c *gin.Context) {
	var input ent.City
	var city ent.City
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	weather, err := h.weatherapi.GetWeather(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.Printf("sending req on openweather.api")
	city, err = h.services.Cities.GetByName(input.Name)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "haven't found this city in table cities. Try to add it first")
		return
	}

	weather.CityId = city.Id
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": weather,
	})
}
