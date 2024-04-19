package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newResponseError(c *gin.Context, status int, message string) {
	logrus.Error(message)

	c.AbortWithStatusJSON(status, Error{Message: message})
}
