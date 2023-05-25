package response

import "time"

type PostResponse struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Status       int       `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	CreateUserId int       `json:"created_user_id"`
}
