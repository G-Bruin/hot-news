package routes

import (
	"github.com/gin-gonic/gin"
	"hotNews/utils"
)

func Init() {
	gin.SetMode(utils.AppSetting.DebugMode)
	router := gin.Default()
	// v1 api
	//v1 := router.Group("/v1")
	//{
	//	v1.GET("/zhihu/top", controllers.ZhTop)
	//}
	port := utils.AppSetting.Port
	router.Run(":" + port)
}
