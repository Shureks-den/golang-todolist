package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	conn "todolist/database"
	"todolist/internal/pkg/middleware"
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

	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	log.Default().Printf("start serving ::%s\n", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("cannot start server on addr 5000:", err.Error())
		return
	}
}
