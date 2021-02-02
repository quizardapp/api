package store

import (
	"github.com/quizardapp/auth-api/internal/model"
)

type UserRepo interface {
	Create(*model.User) error
	Find(value string, field string) (*model.User, error)
	Update(value string, field string, id string) error
}

type CourseRepo interface {
	Create(*model.Course) error
	Find(id string) (*model.Course, error)
	Read(id string) ([]*model.Course, error)
	Update(value string, field string, id string) error
	Delete(id string) error
}

type ModuleRepo interface {
	Create(*model.Module) error
	Find(id string) (*model.Module, error)
	Read(id string) ([]*model.Module, error)
	Update(value string, field string, id string) error
	Delete(id string) error
}

type CardRepo interface {
	Create(*model.Card) error
	Find(id string) (*model.Card, error)
	Read(id string) ([]*model.Card, error)
	Update(value string, field string, id string) error
	Delete(id string) error
}
