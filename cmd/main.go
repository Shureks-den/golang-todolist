package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	conn "todolist/database"
	"todolist/internal/pkg/middleware"
	"todolist/internal/pkg/task/delivery"
	"todolist/internal/pkg/task/repository"
	"todolist/internal/pkg/task/usecase"
)

const port = "5000"

func main() {
	connString := "postgres://todolist:password@localhost:5432/todolist"
	db, err := conn.ConnectToDb(connString)
	if err != nil {
		log.Fatal(err)
	}
	// закрытие коннекта к базе
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("error closing db", err.Error())
		}
	}(db)

	tr := repository.NewTaskRepository(db)
	tu := usecase.NewTaskUsecase(tr)
	td := delivery.NewTaskDelivery(tu)

	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	td.Routing(r)
	log.Default().Printf("start serving ::%s\n", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("cannot start server on addr 5000:", err.Error())
		return
	}
}
