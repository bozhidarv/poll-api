package services

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"github.com/bozhidarv/poll-api/polls/models"
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

func checkDb() error {
	if db == nil {
		return errors.New("DB is null")
	}

	if db == nil {
		err := OpenDbConnection()
		if err != nil {
			return err
		}
	}

	err := db.Ping()
	if err != nil {
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
		return polls, err
	}

	for rows.Next() {
		poll := new(models.Poll)

		err := rows.Scan(&poll.Id, &poll.Name, &poll.Fields, &poll.CreatedBy, &poll.LastUpdated)
		if err != nil {
			return polls, err
		}

		polls = append(polls, *poll)
	}

	return polls, nil
}

func InsertNewPoll(poll models.Poll) error {
	err := checkDb()
	if err != nil {
		return err
	}

	fieldsStr, err := json.Marshal(poll.Fields)
	if err != nil {
		return err
	}
	createdBy := uuid.New().String()

	_, err = db.Exec(`INSERT INTO public.polls
		("name", fields, created_by, last_updated)
		VALUES($1, $2, $3, $4);`,
		poll.Name, fieldsStr, createdBy, time.Now().UTC())
	if err != nil {
		return err
	}

	return nil
}

func UpdatePoll(id string, poll models.Poll) error {
	err := checkDb()
	if err != nil {
		return err
	}

	fieldsStr, err := json.Marshal(poll.Fields)
	if err != nil {
		return err
	}

	_, err = db.Exec(`UPDATE public.polls
		SET "name"=$1, fields=$2, last_updated=$3;
		WHERE id=$4`,
		poll.Name, fieldsStr, time.Now().UTC(), id)
	if err != nil {
		return err
	}

	return nil
}

func CloseDbConn() error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}
