package task

import (
    model "github.com/Jefschlarski/golang-modular/modules/task/model"
)

type TaskRepositoryInterface interface {
    GetTask(id int) (user *model.Task, err error)
    GetTasks() (users []*model.Task, err error)
    CreateTask(createTask model.Task) (userId int, err error)
    UpdateTask(id int, user model.Task) (rowsAffected int64, err error)
    DeleteTask(id int) (rowsAffected int64, err error)
}
