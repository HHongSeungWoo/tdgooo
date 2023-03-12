package model

import (
	"database/sql"
	"database/sql/driver"
	"github.com/goccy/go-json"
	"time"
)

type NullString sql.NullString

//goland:noinspection GoMixedReceiverTypes
func (x *NullString) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(x.String)
}

//goland:noinspection GoMixedReceiverTypes
func (x *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}
	return nil
}

//goland:noinspection GoMixedReceiverTypes
func (x NullString) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	return x.String, nil
}

type Todo struct {
	Id          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `json:"title"`
	Category    NullString `json:"category"`
	Description NullString `json:"description"`
	CreatedAt   time.Time  `gorm:"<-:create" json:"createdAt"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updatedAt"`
}
