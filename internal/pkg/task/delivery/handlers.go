package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"todolist/internal/models"
	"todolist/internal/pkg/task"
)

type TaskDelivery struct {
	taskUsecase task.TaskUsecase
}

func NewTaskDelivery(taskUsecase task.TaskUsecase) *TaskDelivery {
	return &TaskDelivery{
		taskUsecase: taskUsecase,
	}
}

func (td *TaskDelivery) Routing(r *mux.Router) {
	// создание новой задачи
	r.HandleFunc("/task/create", td.CreateTaskHandler).Methods(http.MethodPost, http.MethodOptions)
	// удаление
	r.HandleFunc("/task/{title}", td.DeleteTaskHandler).Methods(http.MethodDelete, http.MethodOptions)
	// получение всех задач
	r.HandleFunc("/tasks", td.SelectAllTaskHandler).Methods(http.MethodGet, http.MethodOptions)
	// получение одной задачи по названию
	r.HandleFunc("/task/{title}", td.DeleteTaskHandler).Methods(http.MethodGet, http.MethodOptions)
	// post запрос будет менять статус
	r.HandleFunc("/task/{title}", td.UpdateTaskHandler).Methods(http.MethodPost, http.MethodOptions)
}

// CreateTaskHandler godoc
// @Summary Creates new task
// @Produce json
// @Param title body string true "Task title"
// @Param description body string false "Task full description"
// @Success 201 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 409 {object} models.Message
// @Router /task/create [post]
func (td *TaskDelivery) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskInput := &models.Task{}
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		buf, _ := json.Marshal(models.Message{Message: "incorrect body"})
		w.Write(buf)
		return
	}
	err = json.Unmarshal(buf, taskInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		buf, _ := json.Marshal(models.Message{Message: "Unable to unmarshal json"})
		w.Write(buf)
		return
	}
	err = td.taskUsecase.CreateTask(taskInput)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		buf, _ := json.Marshal(models.Message{Message: "task already exist"})
		w.Write(buf)
		return
	}
	w.WriteHeader(http.StatusCreated)
	buf, _ = json.Marshal(models.Message{Message: "Task created"})
	w.Write(buf)
}

// DeleteTaskHandler godoc
// @Summary Delete existing task
// @Produce json
// @Param title path string true "Task title"
// @Success 200 {object} models.Message
// @Failure 404 {object} models.Message
// @Router /task/{title} [delete]
func (td *TaskDelivery) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	defer r.Body.Close()
	err := td.taskUsecase.DeleteTask(title)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		buf, _ := json.Marshal(models.Message{Message: "cannot delete task. Task not found"})
		w.Write(buf)
		return
	}
	w.WriteHeader(http.StatusOK)
	buf, _ := json.Marshal(models.Message{Message: "Task deleted"})
	w.Write(buf)
}

// UpdateTaskHandler godoc
// @Summary Change task status
// @Produce json
// @Param title path string true "Task title"
// @Param finished query bool true "Param to change task status"
// @Success 200 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 500 {object} models.Message
// @Router /task/{title} [post]
func (td *TaskDelivery) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	query := r.URL.Query()
	status, err := strconv.ParseBool(query.Get("finished"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		buf, _ := json.Marshal(models.Message{Message: "parsing error"})
		w.Write(buf)
		return
	}
	err = td.taskUsecase.UpdateTask(title, status)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		buf, _ := json.Marshal(models.Message{Message: "cannot update task. Task not found"})
		w.Write(buf)
		return
	}
	w.WriteHeader(http.StatusOK)
	buf, _ := json.Marshal(models.Message{Message: "Task's status updated"})
	w.Write(buf)
}

// SelectAllTaskHandler godoc
// @Summary Get all tasks
// @Produce json
// @Success 200 {object} []models.Task
// @Failure 404 {object} models.Message
// @Failure 500 {object} models.Message
// @Router /tasks [get]
func (td *TaskDelivery) SelectAllTaskHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := td.taskUsecase.GetAllTasks()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		buf, _ := json.Marshal(models.Message{Message: "cannot find tasks"})
		w.Write(buf)
		return
	}
	buf, err := json.Marshal(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		buf, _ := json.Marshal(models.Message{Message: "cannot marshal tasks"})
		w.Write(buf)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}

// GetSingleTaskHandler godoc
// @Summary Get one task by its name
// @Produce json
// @Param title path string true "Task title"
// @Success 200 {object} models.Task
// @Failure 404 {object} models.Message
// @Failure 500 {object} models.Message
// @Router /task/{title} [get]
func (td *TaskDelivery) GetSingleTaskHandler(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	tasks, err := td.taskUsecase.GetSingleTask(title)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		buf, _ := json.Marshal(models.Message{Message: "cannot find task with this title"})
		w.Write(buf)
		return
	}
	buf, err := json.Marshal(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		buf, _ := json.Marshal(models.Message{Message: "cannot marshal tasks"})
		w.Write(buf)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}
