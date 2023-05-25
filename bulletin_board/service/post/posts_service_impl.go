package service

import (
	postinterfaces "gin_test/bulletin_board/dao/post"
	"gin_test/bulletin_board/data/request"
	"gin_test/bulletin_board/data/response"
	"gin_test/bulletin_board/helper"
	"gin_test/bulletin_board/model"

	"github.com/go-playground/validator/v10"
)

func NewPostsRepositoryImpl(postInterface postinterfaces.PostsInterface, validate *validator.Validate) PostsService {
	return &PostsServiceImpl{
		postsInterface: postInterface,
		validate:       validate,
	}
}

type PostsServiceImpl struct {
	postsInterface postinterfaces.PostsInterface
	validate       *validator.Validate
}

// Create implements TagsService
func (t *PostsServiceImpl) Create(tags request.CreatePostsRequest, userId int) error {
	err := t.validate.Struct(tags)
	if err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			return validationErrs
		}
		return err
	}

	tagModel := model.Posts{
		Title:        tags.Title,
		Description:  tags.Description,
		Status:       1,
		CreateUserId: userId,
	}
	err = t.postsInterface.Save(tagModel)
	if err != nil {
		return err
	}

	return nil
}

// Delete implements TagsService
func (t *PostsServiceImpl) Delete(tagsId int) {
	t.postsInterface.Delete(tagsId)
}

// FindAll implements TagsService
func (t *PostsServiceImpl) FindAll() []response.PostResponse {
	result := t.postsInterface.FindAll()
	var tags []response.PostResponse
	for _, value := range result {
		tag := response.PostResponse{
			Id:           value.Id,
			Title:        value.Title,
			Description:  value.Description,
			Status:       value.Status,
			CreatedAt:    value.CreatedAt,
			CreateUserId: value.CreateUserId,
		}
		tags = append(tags, tag)
	}
	return tags
}

// FindById implements TagsService
func (t *PostsServiceImpl) FindById(tagsId int) response.PostResponse {
	tagData, err := t.postsInterface.FindById(tagsId)
	helper.ErrorPanic(err)
	tagResponse := response.PostResponse{
		Id:          tagData.Id,
		Title:       tagData.Title,
		Description: tagData.Description,
		Status:      tagData.Status,
	}
	return tagResponse
}

// Update implements TagsService
func (t *PostsServiceImpl) Update(posts request.UpdatePostsRequest) error {
	postData, err := t.postsInterface.FindById(posts.Id)
	helper.ErrorPanic(err)
	postData.Title = posts.Title
	postData.Description = posts.Description
	postData.Status = *posts.Status
	uerr := t.postsInterface.Update(postData)
	if uerr != nil {
		return err
	}

	return nil
}
