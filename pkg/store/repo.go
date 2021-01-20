package store

import "github.com/quizardapp/auth-api/pkg/model"

type UserRepo interface {
	Create(*model.User) error
	FindByEmail(email string) (*model.User, error)
}
