package interfaces

import (
	"errors"
	"gin_test/bulletin_board/data/request"
	"gin_test/bulletin_board/helper"
	"gin_test/bulletin_board/model"

	"gorm.io/gorm"
)

type PostsRepositoryImpl struct {
	Db *gorm.DB
}

func NewPostsRepositoryImpl(Db *gorm.DB) PostsInterface {
	return &PostsRepositoryImpl{Db: Db}
}

// Delete implements TagsRepository
func (t *PostsRepositoryImpl) Delete(tagsId int) {
	var posts model.Posts
	result := t.Db.Where("id=?", tagsId).Delete(&posts)
	helper.ErrorPanic(result.Error)
}

// FindAll implements TagsRepository
func (t *PostsRepositoryImpl) FindAll() []model.Posts {
	var tags []model.Posts
	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}

// FindById implements TagsRepository
func (t *PostsRepositoryImpl) FindById(tagsId int) (tags model.Posts, err error) {
	var tag model.Posts
	result := t.Db.Find(&tag, tagsId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

// Save implements TagsRepository
func (t *PostsRepositoryImpl) Save(tags model.Posts) error {
	result := t.Db.Create(&tags)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Update implements TagsRepository
func (t *PostsRepositoryImpl) Update(posts model.Posts) error {
	var updateTag = request.UpdatePostsRequest{
		Id:          posts.Id,
		Title:       posts.Title,
		Description: posts.Description,
		Status:      &posts.Status,
	}
	result := t.Db.Model(&posts).Updates(updateTag)
	if result.Error != nil {
		return result.Error
	}
	return nil
	// helper.ErrorPanic(result.Error)
}
