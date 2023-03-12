package model

import (
	"database/sql"
	"time"
)

type TODO struct {
	Id          uint
	Title       string
	Category    sql.NullString
	Description sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
