package module

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	jwt2 "tsqqqqqq/util/jwt"
)

type Users struct {
	Id       int
	Username string
	Password string
	jwt.StandardClaims
}

func (u *Users) GetToken() (res interface{}) {
	token, err := jwt2.GenerateToken(nil, u)
	if err != nil {
		return nil
	}
	fmt.Println(token)
	maps, err := jwt2.ParseToken(token, nil)
	if err != nil {
		return
	}
	fmt.Println(maps)
	return token
}
