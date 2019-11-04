package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TopResult struct {
	Fresh_Text string    `json:"fresh_text"`
	Paging     Paging    `json:"paging"`
	Data       []TopData `json:"data"`
}

type Paging struct {
	Is_End   bool   `json:"is_end"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
}

type TopData struct {
	Detail_Text string     `json:"detail_text"`
	Target      Target     `json:"target"`
	Children    []Children `json:"children"`
}

type Target struct {
	Title string `json:"title"`
	Id    int    `json:"id"`
}

type Children struct {
	Thumbnail string `json:"thumbnail"`
	Type      string `json:"type"`
}

func ZhTop(c *gin.Context) {
	url := "https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total?limit=50&desktop=true"
	body, _ := Curl("GET", url, "")
	var results TopResult
	err := json.Unmarshal([]byte(string(body)), &results)
	if err != nil {
		fmt.Println(err)
	}
	ReturnJson(c, http.StatusOK, "success", results)
}
