package task

import "todolist/internal/models"

type TaskRepository interface {
	AddTask(task *models.Task) error
	DeleteTask(title string) error
	SelectAllTasks() ([]*models.Task, error)
	ChangeTaskStatus(title string, isFinished bool) error
}
