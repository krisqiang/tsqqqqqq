package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const checkAuth = "select rule from permission where permission = ?"

func Permission(c *gin.Context) {
	rule, ok := c.Get("rule")
	if ok {

		rows, err := Query(checkAuth, c.Request.RequestURI)
		if err != nil {
			panic(err)
		}
		var url []string
		if rows.Next() {
			var str string
			err := rows.Scan(&str)
			if err != nil {
				panic(err)
			}
			url = strings.Split(str, ",")
		}
		check := false
		for _, v := range url {
			if v == rule {
				check = true
				break
			}
		}
		if !check {
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "当前登录用户无该接口操作权限",
			})
			c.Abort()
			return
		}
	}

}
