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

func newTaskDelivery(taskUsecase task.TaskUsecase) *TaskDelivery {
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
	// post запрос будет менять статус
	r.HandleFunc("/task/{title}", td.UpdateTaskHandler).Methods(http.MethodPost, http.MethodOptions)
}

func (td *TaskDelivery) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskInput := &models.Task{}
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Incorrect body"))
		return
	}
	err = json.Unmarshal(buf, taskInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to unmarshal json"))
		return
	}
	err = td.taskUsecase.CreateTask(taskInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error writing to db"))
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Task created"))
}

func (td *TaskDelivery) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	defer r.Body.Close()
	err := td.taskUsecase.DeleteTask(title)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("cannot delete task"))
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task deleted"))
}

func (td *TaskDelivery) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	query := r.URL.Query()
	status, err := strconv.ParseBool(query.Get("finished"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("parsing error"))
		w.Write([]byte(err.Error()))
		return
	}
	err = td.taskUsecase.UpdateTask(title, status)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("cannot update task"))
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task's status updated"))
}

func (td *TaskDelivery) SelectAllTaskHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := td.taskUsecase.GetAllTasks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("cannot find tasks"))
		w.Write([]byte(err.Error()))
		return
	}
	buf, err := json.Marshal(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("cannot marshal tasks"))
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}
