package sqlstore

import (
	"fmt"

	"github.com/quizardapp/auth-api/internal/model"
)

type ModuleRepo struct {
	store *SQLStore
}

func (mr *ModuleRepo) Create(m *model.Module) error {

	query := fmt.Sprintf(`
	INSERT INTO modules (id, name, description, iduser, idcourse, creation_date) 
	VALUES ('%s','%s','%s','%s','%s','%s')
	`, m.ID, m.Name, m.Description, m.UserID, m.CourseID, m.CreationDate)

	if _, err := mr.store.db.Exec(query); err != nil {
		return err
	}

	return nil
}

func (mr *ModuleRepo) Find(value string, field string) (*model.Module, error) {
	return nil, nil
}

func (mr *ModuleRepo) Update(value string, field string, id string) error {
	return nil
}

func (mr *ModuleRepo) Delete() error {
	return nil
}
