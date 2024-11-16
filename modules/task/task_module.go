package task

import (
    "database/sql"
    "github.com/gin-gonic/gin"
    "github.com/Jefschlarski/golang-modular/modules"
    controller "github.com/Jefschlarski/golang-modular/modules/task/controller"
    repository "github.com/Jefschlarski/golang-modular/modules/task/repository"
    service "github.com/Jefschlarski/golang-modular/modules/task/service"
    // m "github.com/Jefschlarski/golang-modular/go-task-manager/pkg/middlewares"
)

type TaskModule struct{}

func NewTaskModule() modules.ModuleInterface {
    return &TaskModule{}
}

func (um *TaskModule) Init(router *gin.RouterGroup, db *sql.DB) {
    taskRepository := repository.NewTaskRepository(db)
    taskService := service.NewTaskService(taskRepository)
    taskController := controller.NewTaskController(taskService)

    // middlewares examples
    // router.GET("/tasks", m.Logger(m.Authenticate(taskController.GetTasks)))
    
    router.GET("/tasks", taskController.GetTasks)
    router.GET("/task/:id", taskController.GetTask)
    router.POST("/task", taskController.CreateTask)
    router.PUT("/task/:id", taskController.UpdateTask)
    router.DELETE("/task/:id", taskController.DeleteTask)
}
