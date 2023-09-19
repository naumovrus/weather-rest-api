package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ent "github.com/naumovrus/weather-api/internal/entities"
)

func (h *Handler) GetWeather(c *gin.Context) {
	var input ent.City
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	weather, err := h.weatherapi.GetWeather(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"temp": weather,
	})
}
