package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Poll struct {
	Id          *string    `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Fields      Fields     `json:"fields,omitempty"`
	CreatedBy   *string    `json:"created_by,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
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
