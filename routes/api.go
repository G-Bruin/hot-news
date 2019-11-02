package routes

import (
	"github.com/gin-gonic/gin"
	"hotNews/Http/Controllers"
	"hotNews/utils"
)

func Init() {
	gin.SetMode(utils.AppSetting.DebugMode)
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/login", controllers.Login)
	}
	port := utils.AppSetting.Port
	router.Run(":" + port)
}
