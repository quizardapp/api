package sqlstore

import "github.com/quizardapp/auth-api/internal/model"

type ModuleRepo struct {
	store *SQLStore
}

func (cr *ModuleRepo) Create(m *model.Module) error {
	return nil
}

func (cr *ModuleRepo) Read() error {
	return nil
}

func (cr *ModuleRepo) Update() error {
	return nil
}

func (cr *ModuleRepo) Delete() error {
	return nil
}
