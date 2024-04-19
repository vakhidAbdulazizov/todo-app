package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationToken = "Authorization"
	userKey            = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationToken)

	if header == "" {
		newResponseError(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerData := strings.Split(header, " ")

	if len(headerData) != 2 {
		newResponseError(c, http.StatusUnauthorized, "invalid auth token")
		return
	}

	userId, err := h.service.Authorization.ParseToken(headerData[1])

	if err != nil {
		newResponseError(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userKey, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userKey)

	if !ok {
		newResponseError(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	formatId, ok := id.(int)

	if !ok {
		newResponseError(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}

	return formatId, nil
}
