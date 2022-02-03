package task

import "todolist/internal/models"

type TaskUsecase interface {
	CreateTask(task *models.Task) error
	DeleteTask(title string) error
	UpdateTask(title string, status bool) error
	GetAllTasks() ([]*models.Task, error)
}
