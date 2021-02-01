package teststore

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/quizardapp/auth-api/internal/model"
	"golang.org/x/crypto/bcrypt"
)

// UserRepo ...
type UserRepo struct {
	store *Store
	users map[string]*model.User
}

// Create ...
func (r *UserRepo) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		return errors.New("cannot create uuid")
	}

	u.ID = strings.TrimSuffix(string(uuid), "\n")
	r.users[u.ID] = u

	return nil
}

// Find working only with ID
func (r *UserRepo) Find(value string, field string) (*model.User, error) {

	if r.users[value] == nil {
		return nil, errors.New("user not found")
	}

	return r.users[value], nil

}

// Update ...
func (r *UserRepo) Update(value string, field string, id string) error {

	switch field {
	case "firstname":
		r.users[id].Firstname = value
	case "lastname":
		r.users[id].Lastname = value
	case "email":
		r.users[id].Email = value
	case "password":
		byteValue, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		r.users[id].Password = string(byteValue)
	default:
		return errors.New("unknown field: " + field)
	}

	return nil
}
