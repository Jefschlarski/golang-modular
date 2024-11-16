package task

import (
    "database/sql"
    model "github.com/Jefschlarski/golang-modular/modules/task/model"
)

type TaskRepository struct {
    db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepositoryInterface {
    return &TaskRepository{db: db}
}

func (ur *TaskRepository) GetTask(id int) (user *model.Task, err error) {
    return nil, nil
}

func (ur *TaskRepository) GetTasks() (users []*model.Task, err error) {
    return nil, nil
}

func (ur *TaskRepository) CreateTask(createTask model.Task) (userId int, err error) {
    return 0, nil
}

func (ur *TaskRepository) UpdateTask(id int, user model.Task) (rowsAffected int64, err error) {
    return 0, nil
}

func (ur *TaskRepository) DeleteTask(id int) (rowsAffected int64, err error) {
    return 0, nil
}
