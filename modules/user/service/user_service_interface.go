package user

import (
	dto "github.com/Jefschlarski/golang-modular/modules/user/dto"
	model "github.com/Jefschlarski/golang-modular/modules/user/model"
)

type UserServiceInterface interface {
	GetUser(id int) (user *model.User, err error)
	GetUsers() (users []*model.User, err error)
	CreateUser(userDto dto.CreateUserDto) (user *model.User, err error)
	UpdateUser(id int, userDto dto.UpdateUserDto) (rowsAffected int64, err error)
	DeleteUser(id int) (rowsAffected int64, err error)
	UpdatePassword(id int, updatePasswordDto dto.UpdateUserPasswordDto) (rowsAffected int64, err error)
	Login(loginDto dto.LoginDto) (token string, err error)
}
