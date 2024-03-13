package services

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func OpenDbConnection() error {
	connStr := "postgres://postgres:poll-api@localhost:5432/postgres?sslmode=disable"
	localDb, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	db = localDb
	return nil
}

func CheckDb() (error, *sql.DB) {
	if db == nil {
		err := OpenDbConnection()
		if err != nil {
			return err, nil
		}
		return nil, db
	}

	err := db.Ping()
	if err != nil {
		return err, nil
	}
	return nil, db
}

func CloseDbConn() error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}
