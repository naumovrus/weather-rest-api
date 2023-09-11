package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	usernameCtx         = "username"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
	}
	if headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
	}

	if headerParts[1] == "" {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
	}
	// parse token

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user is not found")
		return 0, errors.New("user is not found")
	}
	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is invalid type")
		return 0, errors.New("user id is invalid type")
	}
	return idInt, nil
}

func getUsername(c *gin.Context) (string, error) {
	username, ok := c.Get(usernameCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user is not found")
		return "", errors.New("user is not found")
	}
	usernameString, ok := username.(string)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user username is invalid type")
		return "", errors.New("user username is invalid type")
	}
	return usernameString, nil
}
