package response

import "time"

type PostResponse struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Status       int       `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreateUserId int       `json:"created_user_id"`
	UpdateUserId int       `json:"updated_user_id"`
	Creator      string    `json:"creator"`
	Updator      string    `json:"updator"`
	IsLoggedIn   bool      `json:"is_logged_in"`
}
