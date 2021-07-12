package service

import "tsqqqqqq/login/module"

type Login interface {
	GetToken(u *module.Users) string
}
