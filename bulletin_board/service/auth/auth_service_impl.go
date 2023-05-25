package service

import (
	"errors"
	"fmt"
	interfaces "gin_test/bulletin_board/dao/user"
	"gin_test/bulletin_board/data/request"
	"gin_test/bulletin_board/helper"
	"gin_test/bulletin_board/model"
	"gin_test/bulletin_board/utils"

	"os"
	"time"

	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	UsersInterface interfaces.UsersInterface
	Validate       *validator.Validate
}

func NewAuthServiceImpl(usersInterface interfaces.UsersInterface, validate *validator.Validate) Authservice {
	return &AuthServiceImpl{
		UsersInterface: usersInterface,
		Validate:       validate,
	}
}

// FindByEmail implements Authservice
func (auth *AuthServiceImpl) FindByEmail(email string) model.User {
	user, _ := auth.UsersInterface.FindByEmail(email)

	return user
}

// Login implements Authservice
func (auth *AuthServiceImpl) Login(users request.LoginRequest) (string, error) {
	tokenExpireInStr := os.Getenv("TOKEN_EXPIRED_IN")
	tokenSecret := os.Getenv("TOKEN_SECRET")

	tokenDuration, err := time.ParseDuration(tokenExpireInStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse TOKEN_EXPIRED_IN: %w", err)
	}

	newUser, err := auth.UsersInterface.FindByEmail(users.Email)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	verifyPassword := utils.VerifyPassword(newUser.Password, users.Password)
	if verifyPassword != nil {
		return "", errors.New("invalid email or password")
	}
	token, err_token := utils.GenerateToken(tokenDuration, newUser.Id, tokenSecret)
	helper.ErrorPanic(err_token)
	return token, nil
}

// Register implements Authservice
func (auth *AuthServiceImpl) Register(users request.CreateUserRequest) {
	hashedPassword, err := utils.HashPassword(users.Password)
	helper.ErrorPanic(err)

	newUser := model.User{
		Username:      users.Username,
		Email:         users.Email,
		Password:      hashedPassword,
		Phone:         users.Phone,
		Address:       users.Address,
		Date_Of_Birth: users.Date_Of_Birth,
		Type:          users.Type,
		Profile_Photo: users.Profile_Photo,
	}
	fmt.Print(newUser)
	auth.UsersInterface.Save(newUser)
}
