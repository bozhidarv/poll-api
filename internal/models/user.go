package models

import "time"

type User struct {
	Id          *string    `json:"id,omitempty"`
	Username    *string    `json:"username,omitempty"`
	Email       *string    `json:"email,omitempty"`
	Password    *string    `json:"password,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}
