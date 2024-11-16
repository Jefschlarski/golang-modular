package user

import (
	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	GetUsers(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
}
