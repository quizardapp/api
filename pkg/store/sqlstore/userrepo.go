package sqlstore

import (
	"errors"
	"fmt"

	"github.com/quizardapp/auth-api/pkg/model"
)

// UserRepo ...
type UserRepo struct {
	store *SQLStore
}

// Create inserts new user into the database and returns error if something went wrong
func (ur *UserRepo) Create(u *model.User) error {

	query := fmt.Sprintf(`
		INSERT INTO users (id, firstname, lastname, email, password, creation_date) 
		VALUES ('%s', '%s', '%s', '%s', '%s', '%s')
		`,
		u.ID, u.Firstname, u.Lastname, u.Email, u.Password, u.CreationDate)

	if _, err := ur.store.db.Exec(query); err != nil {
		return err
	}

	return nil
}

// FindByEmail ...
func (ur *UserRepo) FindByEmail(email string) (*model.User, error) {

	u := model.User{}

	if err := ur.store.db.QueryRow("SELECT * FROM users WHERE email={$1}", email).Scan(&u); err != nil {
		return nil, errors.New("cannot find such user")
	}

	return &u, nil
}
