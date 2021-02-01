package teststore

import (
	"github.com/quizardapp/auth-api/internal/model"
	"github.com/quizardapp/auth-api/internal/store"
)

type Store struct {
	userRepo *UserRepo
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepo {
	if s.userRepo != nil {
		return s.userRepo
	}

	s.userRepo = &UserRepo{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepo
}
