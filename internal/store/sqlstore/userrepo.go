package sqlstore

import (
	"database/sql"
	"fmt"

	"github.com/quizardapp/auth-api/internal/model"
)

// UserRepo ...
type UserRepo struct {
	store *SQLStore
}

// Create inserts new user into the database and returns error if something went wrong
func (ur *UserRepo) Create(u *model.User) error {

	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.GenerateToken("refresh"); err != nil {
		return err
	}

	if err := u.GenerateToken("access"); err != nil {
		return err
	}

	query := fmt.Sprintf(`
		INSERT INTO users (id, firstname, lastname, email, password, creation_date, token) 
		VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s')
						`,
		u.ID, u.Firstname, u.Lastname, u.Email, u.Password, u.CreationDate, u.RefreshToken)
	if _, err := ur.store.db.Exec(query); err != nil {
		return err
	}

	return nil
}

// FindByEmail ...
func (ur *UserRepo) FindByEmail(email string) (*model.User, error) {

	u := model.User{}

	query := fmt.Sprintf("SELECT * FROM users WHERE email='%s'", email)
	token := sql.NullString{}
	if err := ur.store.db.QueryRow(query).Scan(
		&u.ID,
		&u.Firstname,
		&u.Lastname,
		&u.Email,
		&u.Password,
		&u.CreationDate,
		&token); err != nil {
		return nil, err
	}
	u.RefreshToken = token.String

	return &u, nil
}

// FindByID ...
func (ur *UserRepo) FindByID(id string) (*model.User, error) {

	u := model.User{}

	query := fmt.Sprintf("SELECT * FROM users WHERE id='%s'", id)
	token := sql.NullString{}
	if err := ur.store.db.QueryRow(query).Scan(
		&u.ID,
		&u.Firstname,
		&u.Lastname,
		&u.Email,
		&u.Password,
		&u.CreationDate,
		&token); err != nil {
		return nil, err
	}
	u.RefreshToken = token.String

	return &u, nil
}

func (ur *UserRepo) UpdateToken(token string, id string) error {

	query := fmt.Sprintf("UPDATE users SET token='%s' WHERE id='%s'", token, id)
	if err := ur.store.db.QueryRow(query).Scan(); err != nil {
		return err
	}

	return nil
}
