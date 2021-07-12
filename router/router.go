package router

import (
	"github.com/gin-gonic/gin"
	blogControler "tsqqqqqq/blog/controller"
	"tsqqqqqq/login/controller"
	"tsqqqqqq/util"
)

func InitRouter() (res *gin.Engine) {
	res = gin.Default()
	res.Use(util.Cors())
	res.GET("/")
	api := res.Group("/api")
	api.Use(util.JwtVerify, util.Permission)
	blog := api.Group("/blog")
	{
		blog.POST("/login", controller.Login)
		blog.GET("/queryList", blogControler.QueryList)
	}
	return
}
