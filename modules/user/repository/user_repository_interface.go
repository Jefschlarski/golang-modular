package user

import (
	model "github.com/Jefschlarski/golang-modular/modules/user/model"
)

type UserRepositoryInterface interface {
	GetUser(id int) (user *model.User, err error)
	GetUsers() (users []*model.User, err error)
	CreateUser(createUser model.User) (userId int, err error)
	UpdateUser(id int, user model.User) (rowsAffected int64, err error)
	DeleteUser(id int) (rowsAffected int64, err error)
	UpdatePassword(id int, newPassword string) (rowsAffected int64, err error)
	GetUserPassword(id int) (password string, err error)
	GetUserByEmail(email string) (user *model.User, err error)
}
