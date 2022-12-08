package entity

import "time"

type Todo struct {
	ID        int
	Text      string
	Status    string
	Priority  string
	CreatedAt time.Time
}
