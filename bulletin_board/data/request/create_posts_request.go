package request

type CreatePostsRequest struct {
	Title       string `validate:"required,min=5,max=50" json:"name"`
	Description string `validate:"required,min=6,max=255" json:"description"`
	Status      int    `json:"status"`
	UserId      int    `json:"userId" validate:"required"`
}
