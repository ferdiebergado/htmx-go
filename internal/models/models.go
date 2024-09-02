package models

import "time"

type Model struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
}
