package database

import (
	"database/sql"
	_ "github.com/jackc/pgx/stdlib"
	"log"
	"time"
)

func ConnectToDb(connStr string) (*sql.DB, error) {
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("error connecting to db %s", err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("error pinging to db %s", err.Error())
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Minute * 5)
	return db, nil
}
