package service

import (
	interfaces "gin_test/bulletin_board/dao/user"
	"gin_test/bulletin_board/data/request"
	"gin_test/bulletin_board/data/response"
	"gin_test/bulletin_board/helper"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UsersInterface interfaces.UsersInterface
	Validate       *validator.Validate
}

func NewUserServiceImpl(usersInterface interfaces.UsersInterface, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UsersInterface: usersInterface,
		Validate:       validate,
	}
}

// FindAll implements UserService
func (u *UserServiceImpl) FindAll() []response.UserResponse {
	result := u.UsersInterface.FindAll()
	var users []response.UserResponse
	for _, value := range result {
		user := response.UserResponse{
			Id:            value.Id,
			Username:      value.Username,
			Email:         value.Email,
			Password:      value.Password,
			Profile_Photo: value.Profile_Photo,
			Type:          value.Type,
			Phone:         value.Phone,
			Address:       value.Address,
			Date_Of_Birth: value.Date_Of_Birth,
			CreatedAt:     value.CreatedAt,
			UpdatedAt:     value.UpdatedAt,
		}
		users = append(users, user)

	}
	return users
}

// Delete implements UserService
func (u *UserServiceImpl) Delete(userId int) {
	u.UsersInterface.Delete(userId)
}

// FindById implements UserService
func (u *UserServiceImpl) FindById(userId int) response.UserResponse {
	userData, err := u.UsersInterface.FindById(userId)
	helper.ErrorPanic(err)
	userResponse := response.UserResponse{
		Id:            userData.Id,
		Username:      userData.Username,
		Email:         userData.Email,
		Type:          userData.Type,
		Phone:         userData.Phone,
		Address:       userData.Address,
		Date_Of_Birth: userData.Date_Of_Birth,
		Profile_Photo: userData.Profile_Photo,
	}
	return userResponse
}

// Update implements UserService
func (u *UserServiceImpl) Update(users request.UpdateUserRequest) error {
	userData, err := u.UsersInterface.FindById(users.Id)
	helper.ErrorPanic(err)
	userData.Username = users.Username
	userData.Email = users.Email
	userData.Type = users.Type
	userData.Phone = users.Phone
	userData.Address = users.Address
	userData.Date_Of_Birth = users.Date_Of_Birth
	userData.UpdateUserId = users.Updated_User_ID
	userData.Profile_Photo = users.Profile_Photo
	uperr := u.UsersInterface.Update(userData)
	if uperr != nil {
		return err
	}
	return nil

}
