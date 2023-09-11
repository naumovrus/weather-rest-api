package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ent "github.com/naumovrus/weather-api/internal/entities"
)

func (h *Handler) addCity(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input ent.City
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Cities.AddCity(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllCityiesResponse struct {
	Data []ent.City `json:"data"`
}

func (h *Handler) getUsersCity(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	cityies, err := h.services.GetUsersCity(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, getAllCityiesResponse{
		Data: cityies,
	})
}

func (h *Handler) getAllCityies(c *gin.Context) {
	cityies, err := h.services.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, getAllCityiesResponse{
		Data: cityies,
	})
}
