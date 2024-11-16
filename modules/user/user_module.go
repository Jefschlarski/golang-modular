package user

import (
	"database/sql"

	"github.com/Jefschlarski/golang-modular/modules"
	controller "github.com/Jefschlarski/golang-modular/modules/user/controller"
	repository "github.com/Jefschlarski/golang-modular/modules/user/repository"
	service "github.com/Jefschlarski/golang-modular/modules/user/service"
	m "github.com/Jefschlarski/golang-modular/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

type UserModule struct{}

func NewUserModule() modules.ModuleInterface {
	return &UserModule{}
}

func (um *UserModule) Init(router *gin.RouterGroup, db *sql.DB) {
	// Initialize user repository
	userRepository := repository.NewUserRepository(db)

	// Initialize user service
	userService := service.NewUserService(userRepository)

	// Initialize user controller
	userController := controller.NewUserController(userService)

	// Initialize user routes
	router.GET("/users", m.Logger(m.Authenticate(userController.GetUsers)))
	router.GET("/user/:id", m.Logger(m.Authenticate(userController.GetUser)))
	router.POST("/user", m.Logger(userController.CreateUser))
	router.PUT("/user/:id", m.Logger(m.Authenticate(userController.UpdateUser)))
	router.DELETE("/user/:id", m.Logger(m.Authenticate(userController.DeleteUser)))
	router.PATCH("/user/:id", m.Logger(m.Authenticate(userController.UpdatePassword)))
	router.POST("auth/login", m.Logger(userController.Login))
}
