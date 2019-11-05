package controllers

import (
	"encoding/json"
	"fmt"
	"hotNews/cache"
	mysql "hotNews/db"
	"hotNews/http/models"
	"regexp"
	"strconv"
	"time"
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

func ZhTop() {
	limit := 1
	key := "limiter:zhihu-top"
	res, _ := cache.Get(key)
	int_res, _ := strconv.Atoi(string(res))
	if int_res > 0 {
		fmt.Println(key + "数据抓取完成")
		return
	}
	db := mysql.DbEngin
	application := model.Application{}
	//获取app数据 增加计算时间
	appDb := db.Where("alias = ?", key)
	appDb.First(&application)
	if application.Id < 0 {
		return
	}
	number := cache.Limiter(key, limit, application.Polling)
	if !number {
		return
	}

	body, _ := Curl("GET", application.Url, "")
	var result TopResult
	_ = json.Unmarshal([]byte(string(body)), &result)

	article := model.Article{}
	for _, item := range result.Data {
		article.TargetId = strconv.Itoa(item.Target.Id)
		article.ApplicationId = 1
		tmpDb := db.Where("target_id = ?", article.TargetId).Where("application_id = ?", article.ApplicationId)
		tmpDb.First(&article)

		hit, _ := strconv.Atoi(string(regexp.MustCompile("\\d+").Find([]byte(item.Detail_Text))))
		article.Hit = hit * 10000
		article.Title = item.Target.Title
		jsonBytes, _ := json.Marshal(item)
		article.Json = string(jsonBytes)
		if len(item.Children) > 0 {
			article.Cover = item.Children[0].Thumbnail
		}
		if article.Id > 0 {
			tmpDb.Save(&article)
		} else {
			article.CreatedAt = time.Now()
			tmpDb.Create(&article)
		}
		article.Id = 0
	}
	//ReturnJson(c, http.StatusOK, "success", result.Data)
}
