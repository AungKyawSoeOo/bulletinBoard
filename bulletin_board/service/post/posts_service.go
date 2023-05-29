package service

import (
	"gin_test/bulletin_board/data/request"
	"gin_test/bulletin_board/data/response"
)

type PostsService interface {
	Create(tags request.CreatePostsRequest, userId int) error
	Update(tags request.UpdatePostsRequest, userId int) error
	Delete(tagsId int)
	FindById(tagsId int) response.PostResponse
	FindPostByUserId(userId int) []response.PostResponse
	FindAll() []response.PostResponse
}
