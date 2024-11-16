package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	dto "github.com/Jefschlarski/golang-modular/modules/user/dto"
	service "github.com/Jefschlarski/golang-modular/modules/user/service"
)

type UserController struct {
	s service.UserServiceInterface
}

func NewUserController(service service.UserServiceInterface) UserControllerInterface {
	return &UserController{s: service}
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := uc.s.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	returnUserDto := dto.ReturnUserDto{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		UserType:  user.UserType,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, returnUserDto)
}

func (uc *UserController) GetUsers(ctx *gin.Context) {
	users, err := uc.s.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	returnUserDto := []dto.ReturnUserDto{}
	for _, user := range users {
		returnUserDto = append(returnUserDto, dto.ReturnUserDto{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Phone:     user.Phone,
			UserType:  user.UserType,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	ctx.JSON(http.StatusOK, returnUserDto)
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	createUserDto := dto.CreateUserDto{}

	if err := ctx.BindJSON(&createUserDto); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := createUserDto.Validate(); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	user, err := uc.s.CreateUser(createUserDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	returnUserDto := dto.ReturnUserDto{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		UserType:  user.UserType,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	ctx.JSON(http.StatusCreated, returnUserDto)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	updateUserDto := dto.UpdateUserDto{}
	if err := ctx.BindJSON(&updateUserDto); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := updateUserDto.Validate(); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	rowsAffected, err := uc.s.UpdateUser(id, updateUserDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("rows affected %d", rowsAffected))
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := uc.s.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("rows affected %d", rowsAffected))
}

func (uc *UserController) UpdatePassword(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	updateUserPasswordDto := dto.UpdateUserPasswordDto{}
	if err := ctx.BindJSON(&updateUserPasswordDto); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := updateUserPasswordDto.Validate(); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	rowsAffected, err := uc.s.UpdatePassword(id, updateUserPasswordDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("rows affected %d", rowsAffected))
}

func (uc *UserController) Login(ctx *gin.Context) {
	loginDto := dto.LoginDto{}
	if err := ctx.BindJSON(&loginDto); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := uc.s.Login(loginDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func (uc *UserController) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "logout")
}
