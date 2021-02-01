package model

import "time"

// Card ...
type Card struct {
	Name         string
	Content      string
	UserID       string
	ModuleID     string
	CreationDate time.Time
}
