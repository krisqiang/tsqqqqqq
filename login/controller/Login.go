package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	response "tsqqqqqq/base"
	"tsqqqqqq/login/module"
)

func Login(c *gin.Context) {
	res := new(response.Result)
	u := new(module.Users)
	c.BindJSON(u)
	token := u.GetToken()
	res.Success().Put(token)
	fmt.Println(*res)
	//返回自定义结构体需要将结构体的内存地址返回 而不是结构体的引用
	c.JSON(http.StatusOK, res)
}
