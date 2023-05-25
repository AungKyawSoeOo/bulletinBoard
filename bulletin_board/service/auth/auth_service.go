package service

import (
	"gin_test/bulletin_board/data/request"
	"gin_test/bulletin_board/model"
)

type Authservice interface {
	Login(users request.LoginRequest) (string, error)
	Register(users request.CreateUserRequest)
	FindByEmail(email string) model.User
}
