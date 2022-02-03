package usecase

import (
	"log"
	"todolist/internal/models"
	"todolist/internal/pkg/task"
)

type TaskUsecase struct {
	repo   task.TaskRepository
	logger *log.Logger
}

func NewTaskUsecase(repo task.TaskRepository) task.TaskUsecase {
	return &TaskUsecase{
		repo:   repo,
		logger: log.Default(),
	}
}

func (tu *TaskUsecase) CreateTask(task *models.Task) error {
	err := tu.repo.AddTask(task)
	return err
}

func (tu *TaskUsecase) DeleteTask(title string) error {
	err := tu.repo.DeleteTask(title)
	return err
}

func (tu *TaskUsecase) UpdateTask(title string, status bool) error {
	err := tu.repo.ChangeTaskStatus(title, status)
	return err
}

func (tu *TaskUsecase) GetAllTasks() ([]*models.Task, error) {
	tasks, err := tu.repo.SelectAllTasks()
	return tasks, err
}
