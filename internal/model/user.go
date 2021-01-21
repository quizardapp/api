package model

import (
	"errors"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
	ID           string    `json:"id" validate:"required"`
	Firstname    string    `json:"firstname" validate:"required"`
	Lastname     string    `json:"lastname" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	Password     string    `json:"password" validate:"required,gte=6"`
	CreationDate time.Time `json:"creation_date" validate:"required"`
	JWT          string    `json:"jwt"`
}

func (u *User) GenerateJWT() error {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = u.Firstname + " " + u.Lastname
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	if secret := os.Getenv("TOKEN_SECRET"); secret == "" {
		return errors.New("cannot find jwt secret key")
	}

	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return err
	}

	u.JWT = tokenString

	return nil
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
