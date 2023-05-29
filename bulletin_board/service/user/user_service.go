package service

import (
	"gin_test/bulletin_board/data/request"
	"gin_test/bulletin_board/data/response"
)

type UserService interface {
	FindAll() []response.UserResponse
	Delete(userId int)
	FindById(userId int) response.UserResponse
	FindUserById(userId int) []response.UserResponse
	Update(users request.UpdateUserRequest) error
}
