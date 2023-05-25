package request

type CreatePostsRequest struct {
	Title       string `validate:"required,min=1,max=20" json:"name"`
	Description string `validate:"required" json:"description"`
	Status      int    `json:"status"`
	UserId      int    `json:"userId" validate:"required"`
}
