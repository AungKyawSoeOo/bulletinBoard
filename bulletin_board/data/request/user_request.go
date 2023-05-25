package request

import "time"

type CreateUserRequest struct {
	Username        string     `json:"username" validate:"required,min=2,max=100"`
	Email           string     `json:"email" validate:"required,min=2,max=100"`
	Password        string     `json:"password" validate:"required,min=2,max=100"`
	Profile_Photo   string     `json:"profile_photo"`
	Type            string     `default:"0" json:"type"`
	Phone           string     `json:"phone"`
	Address         string     `json:"address"`
	Date_Of_Birth   *time.Time `json:"date_of_birth"`
	Created_User_ID int        `json:"created_user_id"`
}

type UpdateUserRequest struct {
	Id              int        `validate:"required"`
	Username        string     `json:"username" validate:"required,min=2,max=100"`
	Email           string     `json:"email" validate:"required,min=2,max=100"`
	Password        string     `json:"password" validate:"required,min=2,max=100"`
	Profile_Photo   string     `json:"profile_photo"`
	Type            string     `default:"0" json:"type"`
	Phone           string     `json:"phone"`
	Address         string     `json:"address"`
	Date_Of_Birth   *time.Time `json:"date_of_birth"`
	Updated_User_ID int        `json:"updated_user_id"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,min=2,max=100"`
	Password string `json:"password" validate:"required,min=2,max=100"`
}
