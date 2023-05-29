package response

import "time"

type UserResponse struct {
	Id              int        `json:"id"`
	Username        string     `json:"username"`
	Email           string     `json:"email"`
	Password        string     `json:"password"`
	Profile_Photo   string     `json:"profile_photo"`
	Type            string     `default:"0" json:"type"`
	Phone           string     `json:"phone"`
	Address         string     `json:"address"`
	Date_Of_Birth   *time.Time `json:"date_of_birth"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	Created_User_ID int        `json:"created_user_id"`
	Updated_User_ID int        `json:"updated_user_id"`
	Creator         string     `json:"created_user_name"`
	Updator         string     `json:"updated_user_name"`
	// Posts    []PostResponse `json:"posts"`
}

type LoginResponse struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}
