package sqlstore

import (
	"fmt"

	"github.com/quizardapp/auth-api/internal/model"
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

	query := fmt.Sprintf("SELECT * FROM users WHERE email='%s'", email)
	if err := ur.store.db.QueryRow(query).Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.Password, &u.CreationDate); err != nil {
		return nil, err
	}

	return &u, nil
}

// FindByID ...
func (ur *UserRepo) FindByID(id string) (*model.User, error) {

	u := model.User{}

	query := fmt.Sprintf("SELECT * FROM users WHERE id='%s'", id)
	if err := ur.store.db.QueryRow(query).Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.Password, &u.CreationDate); err != nil {
		return nil, err
	}

	return &u, nil
}

// RefreshToken ...
func (ur *UserRepo) RefreshToken(token string) (string, error) {
	return "", nil
}
