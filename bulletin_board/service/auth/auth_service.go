package service

import (
	"gin_test/bulletin_board/data/request"
	"gin_test/bulletin_board/model"
)

type Authservice interface {
	Login(users request.LoginRequest, rememberMe bool) (string, error)
	Register(users request.CreateUserRequest) error
	FindByEmail(email string) model.User
}
