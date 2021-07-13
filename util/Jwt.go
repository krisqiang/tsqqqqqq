package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
	"tsqqqqqq/rule/module"
)

var (
	//自定义的token秘钥
	secret = []byte("563068223")
	//该路由下不校验token
	noVerify = []interface{}{"/api/blog/login", "/", "/api/blog/upload"}
	//token有效时间（纳秒）
	effectTime = 2 * time.Hour
)

// 生成token
func GenerateToken(claims *module.Rule) string {
	//设置token有效期，也可不设置有效期，采用redis的方式
	//   1)将token存储在redis中，设置过期时间，token如没过期，则自动刷新redis过期时间，
	//   2)通过这种方式，可以很方便的为token续期，而且也可以实现长时间不登录的话，强制登录
	//本例只是简单采用 设置token有效期的方式，只是提供了刷新token的方法，并没有做续期处理的逻辑
	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	//生成token
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		//这里因为项目接入了统一异常处理，所以使用panic并不会使程序终止，如不接入，可使用原始方式处理错误
		//接入统一异常可参考 https://blog.csdn.net/u014155085/article/details/106733391
		panic(err)
	}
	return sign
}

//验证token
func JwtVerify(c *gin.Context) {
	//过滤是否验证token
	if IsContainArr(noVerify, c.Request.RequestURI) {
		return
	}
	token := c.GetHeader("token")
	if token == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请求未携带token，无权限访问",
		})
		c.Abort()
		return
	}
	u := parseToken(token)
	if u == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "token 已过期",
		})
		c.Abort()
		return
	}

	//验证token，并存储在请求中
	c.Set("rule", u.Rules)

}

// 解析Token
func parseToken(tokenString string) *module.Rule {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &module.Rule{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil
	}
	claims, ok := token.Claims.(*module.Rule)
	if !ok {
		panic("token is valid")
	}
	return claims
}

// 更新token
func Refresh(tokenString string) string {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &module.Rule{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*module.Rule)
	if !ok {
		panic("token is valid")
	}
	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
	return GenerateToken(claims)
}

func IsContainArr(noVerify []interface{}, requestURI string) bool {
	fmt.Println(noVerify, requestURI)
	res := false
	for _, v := range noVerify {
		fmt.Println(v, requestURI, requestURI == v)
		if v == requestURI {
			res = true
			break
		}
	}
	fmt.Println(res)
	return res
}
