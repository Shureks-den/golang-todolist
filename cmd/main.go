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

	httpSwagger "github.com/swaggo/http-swagger"
	_ "todolist/docs" // docs is generated by Swag CLI, you have to import it.
)

const port = ":5000"

// @title ToDoList api
// @version 1.0
// @description Test Task for building api for todolist
// @termsOfService http://swagger.io/terms/

// @contact.email alexander.klonov@mail.ru

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api
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
	r = r.PathPrefix("/api").Subrouter()
	r.Use(middleware.ContentTypeMiddleware)
	td.Routing(r)

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
	log.Default().Printf("start serving %s\n", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("cannot start server on addr 5000\n", err.Error())
		return
	}
}
