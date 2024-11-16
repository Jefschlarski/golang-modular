package user

import (
	"errors"
	"time"

	dto "github.com/Jefschlarski/golang-modular/modules/user/dto"
	model "github.com/Jefschlarski/golang-modular/modules/user/model"
	repository "github.com/Jefschlarski/golang-modular/modules/user/repository"
	"github.com/Jefschlarski/golang-modular/pkg/security"
)

type UserService struct {
	r repository.UserRepositoryInterface
}

func NewUserService(repository repository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{r: repository}
}

func (us *UserService) GetUser(id int) (user *model.User, err error) {
	user, err = us.r.GetUser(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (us *UserService) GetUsers() (users []*model.User, err error) {
	users, err = us.r.GetUsers()
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("no users found")
	}

	return users, nil
}

func (us *UserService) CreateUser(userDto dto.CreateUserDto) (user *model.User, err error) {

	hashedPassword, err := security.Hash(userDto.Password)
	if err != nil {
		return nil, err
	}

	user = &model.User{
		Name:      userDto.Name,
		Email:     userDto.Email,
		Phone:     userDto.Phone,
		UserType:  1,
		Password:  string(hashedPassword),
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	userId, err := us.r.CreateUser(*user)
	if err != nil {
		return nil, err
	}

	user.ID = userId

	return user, nil
}

func (us *UserService) UpdateUser(id int, userDto dto.UpdateUserDto) (rowsAffected int64, err error) {
	user, err := us.r.GetUser(id)
	if err != nil {
		return 0, err
	}

	if user == nil {
		return 0, errors.New("user not found")
	}

	if userDto.Name != "" {
		user.Name = userDto.Name
	}

	if userDto.Email != "" {
		user.Email = userDto.Email
	}

	if userDto.Phone != "" {
		user.Phone = userDto.Phone
	}

	if userDto.UserType != 0 {
		user.UserType = userDto.UserType
	}

	user.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	rowsAffected, err = us.r.UpdateUser(id, *user)
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (us *UserService) DeleteUser(id int) (rowsAffected int64, err error) {

	rowsAffected, err = us.r.DeleteUser(id)
	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		return 0, errors.New("user not found")
	}

	return rowsAffected, nil
}

func (us *UserService) UpdatePassword(id int, updatePasswordDto dto.UpdateUserPasswordDto) (rowsAffected int64, err error) {
	user_password, err := us.r.GetUserPassword(id)
	if err != nil {
		return 0, err
	}

	if user_password == "" {
		return 0, errors.New("user not found")
	}

	err = security.Compare(user_password, updatePasswordDto.OldPassword)
	if err != nil {
		return 0, err
	}

	newHashedPassword, err := security.Hash(updatePasswordDto.NewPassword)
	if err != nil {
		return 0, err
	}

	return us.r.UpdatePassword(id, string(newHashedPassword))
}

func (us *UserService) Login(loginDto dto.LoginDto) (token string, err error) {

	user, err := us.r.GetUserByEmail(loginDto.Email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("user not found")
	}

	err = security.Compare(user.Password, loginDto.Password)
	if err != nil {
		return "", err
	}

	token, err = security.GenerateToken(uint64(user.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}
