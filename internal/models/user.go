package models

import "time"

type User struct {
	Id          *string    `json:"id"`
	Username    *string    `json:"username"`
	Email       *string    `json:"email"`
	Password    *string    `json:"password"`
	LastUpdated *time.Time `json:"last_updated"`
}
