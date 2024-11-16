package task

import (
   dto "github.com/Jefschlarski/golang-modular/modules/task/dto"
   model "github.com/Jefschlarski/golang-modular/modules/task/model"
)

type TaskServiceInterface interface {
    GetTask(id int) (user *model.Task, err error)
    GetTasks() (users []*model.Task, err error)
    CreateTask(userDto dto.CreateTaskDto) (user *model.Task, err error)
    UpdateTask(id int, userDto dto.UpdateTaskDto) (rowsAffected int64, err error)
    DeleteTask(id int) (rowsAffected int64, err error)
}