package task

import (
    "github.com/gin-gonic/gin"
    service "github.com/Jefschlarski/golang-modular/modules/task/service"
)

type TaskController struct {
    s service.TaskServiceInterface
}

func NewTaskController(service service.TaskServiceInterface) TaskControllerInterface {
    return &TaskController{s: service}
}

func (uc *TaskController) GetTasks(ctx *gin.Context) {
    // Implementação do método GetTasks
}

func (uc *TaskController) GetTask(ctx *gin.Context) {
    // Implementação do método GetTask
}

func (uc *TaskController) CreateTask(ctx *gin.Context) {
    // Implementação do método CreateTask
}

func (uc *TaskController) UpdateTask(ctx *gin.Context) {
    // Implementação do método UpdateTask
}

func (uc *TaskController) DeleteTask(ctx *gin.Context) {
    // Implementação do método DeleteTask
}
