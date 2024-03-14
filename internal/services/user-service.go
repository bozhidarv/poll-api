package services

import (
	"time"

	"github.com/bozhidarv/poll-api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RegisterUser(user models.User) error {
	err, db := CheckDb()
	if err != nil {
		return err
	}

	hashedPass, err := hashPassword(*user.Password)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		"INSERT INTO users (username, email, password, last_updated) VALUES ($1, $2, $3, $4)",
		user.Username,
		user.Email,
		hashedPass,
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}
