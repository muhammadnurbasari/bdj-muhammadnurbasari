package usersModel

import "time"

type Users struct {
	ID        int
	Name      string
	Position  string
	JobDesk   string
	CreatedAt time.Time
}
