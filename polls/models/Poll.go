package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Poll struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Fields      Fields    `json:"fields"`
	CreatedBy   string    `json:"created_by"`
	LastUpdated time.Time `json:"last_updated"`
}

type (
	Field  map[string]interface{}
	Fields []Field
)

func (a Fields) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (ls *Fields) Scan(src any) error {
	var data []byte
	switch v := src.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	}
	return json.Unmarshal(data, ls)
}
