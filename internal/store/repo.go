package store

import "github.com/quizardapp/auth-api/internal/model"

type UserRepo interface {
	Create(*model.User) error
	FindByEmail(email string) (*model.User, error)
	FindByID(id string) (*model.User, error)
	RefreshToken(token string) (string, error)
}
