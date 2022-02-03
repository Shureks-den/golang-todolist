package repository

import (
	"database/sql"
	"log"
	"time"
	"todolist/internal/models"
	"todolist/internal/pkg/task"
)

type TaskRepository struct {
	db     *sql.DB
	logger *log.Logger
}

func NewTaskRepository(db *sql.DB) task.TaskRepository {
	return &TaskRepository{
		db:     db,
		logger: log.Default(),
	}
}

func (tr *TaskRepository) AddTask(task *models.Task) error {
	curTime := time.Now()
	_, err := tr.db.Exec(
		"INSERT INTO tasks (title, description, created) VALUES ($1, $2, $3);",
		task.Title, task.Description, curTime)
	if err != nil {
		tr.logger.Print(err.Error())
		return err
	}
	return nil
}

func (tr *TaskRepository) DeleteTask(title string) error {
	_, err := tr.db.Exec(
		"DELETE FROM tasks WHERE title = $1", title)
	if err != nil {
		tr.logger.Print(err.Error())
		return err
	}
	return nil
}

func (tr *TaskRepository) SelectAllTasks() ([]*models.Task, error) {
	rows, err := tr.db.Query("SELECT id, title, description, isFinished, created FROM tasks")
	if err != nil {
		tr.logger.Print(err.Error())
		return nil, err
	}
	tasks := make([]*models.Task, 0)

	for rows.Next() {
		res := &models.Task{}
		err = rows.Scan(&res.Id, &res.Title, &res.Description, &res.IsFinished, &res.Created)
		if err != nil {
			tr.logger.Print(err.Error())
			return nil, err
		}
		tasks = append(tasks, res)
	}
	return tasks, nil
}

func (tr *TaskRepository) ChangeTaskStatus(title string, isFinished bool) error {
	_, err := tr.db.Exec("UPDATE tasks SET isFinished = $1 WHERE title = $2",
		isFinished, title)
	if err != nil {
		tr.logger.Print(err.Error())
		return err
	}
	return nil
}
