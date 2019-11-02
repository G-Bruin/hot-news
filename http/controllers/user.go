package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {

	url := "https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total"

	list, _ := Curl("GET", url, "limit=50&desktop=true")

	ReturnJson(c, http.StatusOK, "success", list)
}
