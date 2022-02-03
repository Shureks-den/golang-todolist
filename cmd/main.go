package main

import (
	"database/sql"
	"log"
	conn "todolist/database"

	"github.com/gorilla/mux"
	"net/http"
)

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
	err = http.ListenAndServe("5000", r)
	if err != nil {
		log.Fatal("cannot start server on addr 5000:", err.Error())
		return
	}
}
