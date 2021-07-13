package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	blogControler "tsqqqqqq/blog/controller"
	"tsqqqqqq/login/controller"
	"tsqqqqqq/util"
)

func InitRouter() (res *gin.Engine) {
	res = gin.Default()
	res.StaticFS("/image", http.Dir(util.StaticFileRouter()))
	res.MaxMultipartMemory = 8 << 20 // 8 MiB
	res.Use(util.Cors())
	res.GET("/")
	api := res.Group("/api")
	api.Use(util.JwtVerify, util.Permission)
	blog := api.Group("/blog")
	{
		blog.POST("/login", controller.Login)
		blog.GET("/queryList", blogControler.QueryList)
		blog.POST("/upload", blogControler.Upload)
		blog.POST("/save", blogControler.Save)
	}
	return
}
