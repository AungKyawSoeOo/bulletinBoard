package interfaces

import "gin_test/bulletin_board/model"

type UsersInterface interface {
	Save(users model.User) error
	Update(users model.User) error
	Delete(userId int)
	FindById(userId int) (model.User, error)
	FindAll() []model.User
	FindByEmail(email string) (model.User, error)
	FindByUsername(username string) (model.User, error)
	FindUserById(userId int) (users []model.User, err error)
	UpdatePassword(users model.User) error
}
