package router

import (
	"github.com/gin-gonic/gin"
	"tsqqqqqq/login/controller"
)

func InitRouter() (res *gin.Engine) {
	res = gin.Default()
	res.GET("/")
	api := res.Group("/api")
	blog := api.Group("/blog")
	{
		blog.POST("/login", controller.Login)
	}
	return
}
