// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID    `json:"id"`
	Username  string       `json:"username"`
	Password  string       `json:"password"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}
