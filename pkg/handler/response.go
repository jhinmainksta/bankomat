package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResp struct {
	Message string
}

type statusResp struct {
	Status string
}

func errorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResp{Message: message})
}

func succesResponse(c *gin.Context, statusCode int, message string, body any) {
	logrus.Info(message)
	c.JSON(statusCode, body)
}
