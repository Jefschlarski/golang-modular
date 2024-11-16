package task

import (
    dto "github.com/Jefschlarski/golang-modular/modules/task/dto"
    model "github.com/Jefschlarski/golang-modular/modules/task/model"
    repository "github.com/Jefschlarski/golang-modular/modules/task/repository"
)

type TaskService struct {
    r repository.TaskRepositoryInterface
}

func NewTaskService(repository repository.TaskRepositoryInterface) TaskServiceInterface {
    return &TaskService{r: repository}
}

func (us *TaskService) GetTask(id int) (user *model.Task, err error) {
    return nil, nil
}

func (us *TaskService) GetTasks() (users []*model.Task, err error) {
    return nil, nil
}

func (us *TaskService) CreateTask(userDto dto.CreateTaskDto) (user *model.Task, err error) {
    return nil, nil
}

func (us *TaskService) UpdateTask(id int, userDto dto.UpdateTaskDto) (rowsAffected int64, err error) {
    return 1, nil
}

func (us *TaskService) DeleteTask(id int) (rowsAffected int64, err error) {
	return 1, nil
}