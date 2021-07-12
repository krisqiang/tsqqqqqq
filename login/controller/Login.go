package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	response "tsqqqqqq/base"
	"tsqqqqqq/login/module"
	"tsqqqqqq/login/service"
	"tsqqqqqq/login/service/impl"
)

func Login(c *gin.Context) {
	res := new(response.Result)
	u := new(module.Users)
	c.BindJSON(u)
	var loginService service.Login
	loginService = new(impl.LoginImpl)
	token := loginService.GetToken(u)
	res.Success().Put(token)
	//返回自定义结构体需要将结构体的内存地址返回 而不是结构体的引用
	c.JSON(http.StatusOK, res)
}
