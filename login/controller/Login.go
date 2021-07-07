package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tsqqqqqq/login/module"
)

func Login(c *gin.Context) {
	u := new(module.Users)
	c.BindJSON(u)
	token := u.GetToken()

	c.JSON(http.StatusOK, gin.H{"token": token})
}
