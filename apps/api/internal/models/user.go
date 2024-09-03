package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id"`
	Username  string    `db:"username"`
	Color     string    `db:"color"`
	CreatedAt time.Time `db:"created_at"`
}
