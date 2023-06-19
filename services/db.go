package services

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/bozhidarv/poll-api/polls/models"
)

var db *sql.DB

func OpenDbConnection() error {
	connStr := "postgres://postgres:poll-api@localhost:5432/postgres?sslmode=disable"
	localDb, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	db = localDb
	return nil
}

func checkDb() error {
	err := db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func GetAllPolls() ([]models.Poll, error) {
	var polls []models.Poll

	err := checkDb()
	if err != nil {
		return polls, err
	}

	rows, err := db.Query(`SELECT id, name, fields, created_by, last_updated FROM public.polls`)
	if err != nil {
		fmt.Println(err.Error())
		return polls, err
	}

	for rows.Next() {
		poll := new(models.Poll)

		err := rows.Scan(&poll.Id, &poll.Name, &poll.Fields, &poll.CreatedBy, &poll.LastUpdated)
		if err != nil {
			fmt.Println(err.Error())
			return polls, err
		}

		polls = append(polls, *poll)
	}

	return polls, nil
}

func CloseDbConn() error {
	err := db.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
