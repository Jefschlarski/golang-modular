package task

import (
    "github.com/gin-gonic/gin"
)

type TaskControllerInterface interface {
    GetTasks(ctx *gin.Context)
    GetTask(ctx *gin.Context)
    CreateTask(ctx *gin.Context)
    UpdateTask(ctx *gin.Context)
    DeleteTask(ctx *gin.Context)
}