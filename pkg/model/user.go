package model

import "time"

type User struct {
	ID           string    `json:"id"`
	Firstname    string    `json:"firstname"`
	Lastname     string    `json:"lastname"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	CreationDate time.Time `json:"creation_date"`
}
