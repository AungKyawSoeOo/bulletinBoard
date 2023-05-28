package request

import "time"

type UpdatePostsRequest struct {
	Id           int    `validate:"required"`
	Title        string `validate:"required,min=1,max=20" json:"name"`
	Description  string `validate:"required" json:"description"`
	Status       *int   `json:"status"`
	UpdateUserId int    `json:"updated_user_id"`
	UpdatedAt    time.Time
}
