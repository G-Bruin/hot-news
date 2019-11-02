package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	list := make(map[string]interface{})
	list["message"] = c.PostForm("message")
	list["valuename"] = c.DefaultPostForm("valuename", "Golang")
	ReturnJson(c, http.StatusOK, "success", list)
}
