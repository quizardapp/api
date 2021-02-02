package sqlstore

import (
	"fmt"

	"github.com/quizardapp/auth-api/internal/model"
)

// ModuleRepo ...
type ModuleRepo struct {
	store *SQLStore
}

// Create ...
func (mr *ModuleRepo) Create(m *model.Module) error {

	query := fmt.Sprintf(`
	INSERT INTO modules (id, name, description, iduser, creation_date) 
	VALUES ('%s','%s','%s','%s','%s')
	`, m.ID, m.Name, m.Description, m.UserID, m.CreationDate)

	if _, err := mr.store.db.Exec(query); err != nil {
		return err
	}

	return nil
}

// Find ...
func (mr *ModuleRepo) Find(id string) (*model.Module, error) {

	m := model.Module{}

	query := fmt.Sprintf(`SELECT * FROM courses WHERE id='%s'`, id)
	if err := mr.store.db.QueryRow(query).Scan(&m.ID, &m.Name, &m.Description, &m.UserID, &m.CreationDate); err != nil {
		return nil, err
	}

	return &m, nil
}

// Read ...
func (mr *ModuleRepo) Read(id string) ([]*model.Module, error) {

	query := fmt.Sprintf(`SELECT * FROM modules WHERE iduser='%s'`, id)
	rows, err := mr.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	numOfRows := 0
	query = fmt.Sprintf(`SELECT count(id) FROM modules WHERE iduser='%s'`, id)
	if err := mr.store.db.QueryRow(query).Scan(&numOfRows); err != nil {
		return nil, err
	}

	modules := make([]*model.Module, numOfRows)
	i := 0
	for rows.Next() {
		module := model.Module{}

		err = rows.Scan(&module.ID, &module.Name, &module.Description, &module.UserID, &module.CreationDate)
		if err != nil {
			return nil, err
		}
		modules[i] = &module
		i++
	}
	return modules, nil
}

// Update ...
func (mr *ModuleRepo) Update(value string, field string, id string) error {

	query := fmt.Sprintf("UPDATE modules SET `%s`='%s' WHERE id='%s'", field, value, id)
	fmt.Println(query)
	if _, err := mr.store.db.Exec(query); err != nil {
		return err
	}

	return nil
}

// Delete ...
func (mr *ModuleRepo) Delete(id string) error {
	query := fmt.Sprintf("DELETE FROM modules WHERE id='%s'", id)
	if _, err := mr.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}
