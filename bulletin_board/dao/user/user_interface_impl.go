package interfaces

import (
	"errors"
	"fmt"

	"gin_test/bulletin_board/data/request"
	"gin_test/bulletin_board/helper"
	"gin_test/bulletin_board/model"

	"gorm.io/gorm"
)

type UsersInterfaceImpl struct {
	Db *gorm.DB
}

func NewUsersInterfaceImpl(Db *gorm.DB) UsersInterface {
	return &UsersInterfaceImpl{Db: Db}
}

func (u *UsersInterfaceImpl) FindUserById(userId int) ([]model.User, error) {
	var users []model.User
	result := u.Db.Find(&users, "create_user_id = ?", userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// Delete implements UsersInterface
func (user *UsersInterfaceImpl) Delete(userId int) {

	var postModel model.Posts
	presult := user.Db.Where("create_user_id = ?", userId).Delete(&postModel)
	helper.ErrorPanic(presult.Error)
	var users model.User
	result := user.Db.Where("id = ?", userId).Delete(&users)
	helper.ErrorPanic(result.Error)

	// result := userDao.DB.Unscoped().Model(&user).Association("Post").Unscoped().Clear()
	// helper.ErrorPanic(result)
	// userDao.DB.Unscoped().Delete(&user)
}

// FindAll implements UsersInterface
func (user *UsersInterfaceImpl) FindAll() []model.User {
	var users []model.User
	result := user.Db.Preload("Posts").Find(&users)
	// result := user.Db.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

// FindByEmail implements UsersInterface
// func (user *UsersInterfaceImpl) FindByEmail(email string) (model.User, error) {
// 	var users model.User
// 	result := user.Db.Preload("Posts").First(&users, userId)
// 	if result.Error != nil {
// 		return users, errors.New("user not found")
// 	}
// 	return users, nil
// }

// FindById implements UsersInterface
func (u *UsersInterfaceImpl) FindById(userId int) (use model.User, err error) {
	var user model.User
	result := u.Db.Find(&user, userId)
	if result != nil {
		return user, nil
	}
	return user, errors.New("user is not found")
}

// FindByUsername implements UsersInterface
func (user *UsersInterfaceImpl) FindByUsername(username string) (model.User, error) {
	var users model.User
	result := user.Db.First(&users, "username = ?", username)
	if result.Error != nil {
		return users, errors.New("invalid username or password")
	}
	return users, nil
}

// FindByEmail implements UsersInterface.
func (user *UsersInterfaceImpl) FindByEmail(email string) (model.User, error) {
	var users model.User
	result := user.Db.First(&users, "email = ?", email)
	if result.Error != nil {
		return users, errors.New("invalid email or password")
	}
	return users, nil
}

// Save implements UsersInterface
func (user *UsersInterfaceImpl) Save(users model.User) error {
	result := user.Db.Create(&users)
	fmt.Print(result)
	helper.ErrorPanic(result.Error)
	if result.Error != nil {
		return errors.New("something wrong")
	}
	return nil
}

// Update implements UsersInterface
func (user *UsersInterfaceImpl) Update(users model.User) error {
	var updateUsers = request.UpdateUserRequest{
		Id:            users.Id,
		Username:      users.Username,
		Email:         users.Email,
		Password:      users.Password,
		Type:          users.Type,
		Phone:         users.Phone,
		Address:       users.Address,
		Date_Of_Birth: users.Date_Of_Birth,
		UpdateUserId:  users.UpdateUserId,
		Profile_Photo: users.Profile_Photo,
		UpdatedAt:     users.UpdatedAt,
	}

	result := user.Db.Model(&users).Updates(updateUsers)
	helper.ErrorPanic(result.Error)
	if result.Error != nil {
		return errors.New("something wrong")
	}
	return nil
}
