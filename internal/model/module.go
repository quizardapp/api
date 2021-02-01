package model

import "time"

// Module ...
type Module struct {
	ID           string
	Name         string
	Description  string
	UserID       string
	CourseID     string
	CreationDate time.Time
}
