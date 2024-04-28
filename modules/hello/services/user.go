package services

import (
	"gin-demo/base"
)

type UserService struct {
	*base.BaseService
}

var SerUser = UserService{
	base.NewService("hello"),
}
