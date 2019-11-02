package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context)  {
	message := c.PostForm("message")
	valuename := c.DefaultPostForm("valuename", "Golang")

	c.JSON(http.StatusOK, gin.H{"status": gin.H{"status_code": http.StatusOK, "status": "ok",}, "message": message, "name": valuename,})
}