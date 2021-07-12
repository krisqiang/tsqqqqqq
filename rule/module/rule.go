package module

import "github.com/golang-jwt/jwt"

type Rule struct {
	Id      int    //id
	Rules   string //权限名称
	Created string //创建时间
	Updated string //修改时间
	jwt.StandardClaims
}
