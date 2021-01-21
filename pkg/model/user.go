package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string    `json:"id" validate:"required"`
	Firstname    string    `json:"firstname" validate:"required"`
	Lastname     string    `json:"lastname" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	Password     string    `json:"password" validate:"required,gte=6"`
	CreationDate time.Time `json:"creation_date" validate:"required"`
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func (u *User) Validate() error {
	if err := validator.New().Struct(u); err != nil {
		return err
	}
	return nil
}
