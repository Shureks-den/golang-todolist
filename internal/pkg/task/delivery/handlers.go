package delivery

import (
	"github.com/gorilla/mux"
	"net/http"
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
	r.HandleFunc("/task/create")  // создание новой задачи
	r.HandleFunc("/task/{title}") // удаление
	r.HandleFunc("/tasks")        // получение всех задач
	r.Handle("/task/{title}")     // post запрос будет менять статус
}

func (td *TaskDelivery) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {

}
