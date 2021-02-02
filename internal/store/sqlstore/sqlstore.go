package sqlstore

import (
	"database/sql"

	"github.com/quizardapp/auth-api/internal/store"
)

// SQLStore ...
type SQLStore struct {
	db         *sql.DB
	userRepo   *UserRepo
	courseRepo *CourseRepo
	moduleRepo *ModuleRepo
	// cardRepo   *CardRepo
}

// New ...
func New(db *sql.DB) *SQLStore {
	return &SQLStore{
		db: db,
	}
}

// User ...
func (s *SQLStore) User() store.UserRepo {
	if s.userRepo != nil {
		return s.userRepo
	}

	s.userRepo = &UserRepo{store: s}

	return s.userRepo
}

// Course ...
func (s *SQLStore) Course() store.CourseRepo {
	if s.courseRepo != nil {
		return s.courseRepo
	}

	s.courseRepo = &CourseRepo{store: s}

	return s.courseRepo
}

// Module ...
func (s *SQLStore) Module() store.ModuleRepo {
	if s.moduleRepo != nil {
		return s.moduleRepo
	}

	s.moduleRepo = &ModuleRepo{store: s}

	return s.moduleRepo
}

// func (s *SQLStore) Card() store.CardRepo {
// 	if s.cardRepo != nil {
// 		return s.cardRepo
// 	}

// 	s.userRepo = &CardRepo{store: s}

// 	return s.cardRepo
// }
