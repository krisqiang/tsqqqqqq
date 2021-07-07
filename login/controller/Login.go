package controller

import (
	"github.com/gin-gonic/gin"
	"tsqqqqqq/login/module"
)

func Login(c *gin.Context) {
	u := new(module.Users)
	c.BindJSON(u)
	c.JSON(200, gin.H{"users": u})
}
