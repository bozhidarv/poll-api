package services

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/bozhidarv/poll-api/internal/models"
	"github.com/google/uuid"
)

func GetAllPolls() ([]models.Poll, error) {
	polls := make([]models.Poll, 0)
	err, db := CheckDb()
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

func GetPollById(pollUuid string) (models.Poll, error) {
	poll := new(models.Poll)

	err, db := CheckDb()

	if err != nil {
		return *poll, err
	}

	rows, err := db.Query(
		`SELECT id, name, fields, created_by, last_updated FROM public.polls WHERE id = $1`,
		pollUuid,
	)
	if err != nil {
		return *poll, err
	}

	if !rows.Next() {
		return *poll, errors.New("NOT_FOUND")
	}

	err = rows.Scan(&poll.Id, &poll.Name, &poll.Fields, &poll.CreatedBy, &poll.LastUpdated)
	if err != nil {
		return *poll, err
	}

	return *poll, nil
}

func InsertNewPoll(poll models.Poll) error {
	err, db := CheckDb()
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
	err, db := CheckDb()
	if err != nil {
		return err
	}

	fieldsStr, err := json.Marshal(poll.Fields)
	if err != nil {
		return err
	}

	res, err := db.Exec(`UPDATE public.polls
		SET "name"=$1, fields=$2, last_updated=$3
		WHERE id=$4`,
		poll.Name, fieldsStr, time.Now().UTC(), id)
	if err != nil {
		return err
	}

	if n, err := res.RowsAffected(); n == 0 {
		if err != nil {
			return err
		}
		return errors.New("NOT_FOUND")
	}

	return nil
}

func DeletePoll(id string) error {
	err, db := CheckDb()
	if err != nil {
		return err
	}

	res, err := db.Exec(`DELETE FROM public.polls	WHERE id=$1`, id)
	if err != nil {
		return err
	}

	if n, err := res.RowsAffected(); n == 0 {
		if err != nil {
			return err
		}
		return errors.New("NOT_FOUND")
	}

	return nil
}
