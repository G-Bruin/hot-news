package main

import (
	"hotNews/cache"
	"hotNews/db"
	"hotNews/routes"
)

func main() {
	defer db.DbClose()
	cache.Ping()
	// 初始化路由
	routes.Init()
}