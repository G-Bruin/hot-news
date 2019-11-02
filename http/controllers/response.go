package controllers

import (
	"github.com/gin-gonic/gin"
)

func ReturnJson(c *gin.Context, status_code int, message string, data interface{}) {
	c.JSON(
		status_code,
		gin.H{
			"status_code": status_code,
			"message":     message,
			"data":        data,
		})
}
