package sqlstore

import (
	"database/sql"

	"github.com/quizardapp/auth-api/pkg/store"
)

type SQLStore struct {
	db       *sql.DB
	userRepo *UserRepo
}

func New(db *sql.DB) *SQLStore {
	return &SQLStore{
		db: db,
	}
}

func (s *SQLStore) User() store.UserRepo {
	if s.userRepo != nil {
		return s.userRepo
	}

	s.userRepo = &UserRepo{store: s}

	return s.userRepo
}
