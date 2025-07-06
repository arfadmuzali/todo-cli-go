package storage

import "time"

type Todo struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
}
