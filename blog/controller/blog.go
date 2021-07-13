package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	response "tsqqqqqq/base"
	"tsqqqqqq/blog/module"
	"tsqqqqqq/blog/service/impl"
	"tsqqqqqq/util"
)

var blogService = new(impl.BlogServiceImpl)

func QueryList(c *gin.Context) {
	res := new(response.Result)
	res.Success()
	c.JSON(http.StatusOK, res)
}

func Upload(c *gin.Context) {
	res := new(response.Result)
	res.Success()
	file, err := c.FormFile("file")
	fmt.Println(err)
	// 上传文件至指定目录
	local := util.Upload(file)
	fmt.Println("local", local)
	c.JSON(http.StatusOK, res.Put(local))
}

func Save(c *gin.Context) {
	res := new(response.Result)
	res.Success()
	blog := new(module.Blog)
	c.BindJSON(blog)
	b := blogService.Save(blog)
	c.JSON(http.StatusOK, res.Put(b))

}
