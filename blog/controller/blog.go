package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	response "tsqqqqqq/base"
)

func QueryList(c *gin.Context) {
	res := new(response.Result)
	res.Success()
	c.JSON(http.StatusOK, res)
}
