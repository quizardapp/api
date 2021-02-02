package sqlstore

import (
	"fmt"

	"github.com/quizardapp/auth-api/internal/model"
)

// CourseRepo ...
type CourseRepo struct {
	store *SQLStore
}

// Create ...
func (cr *CourseRepo) Create(c *model.Course) error {

	if err := c.Validate(); err != nil {
		return err
	}

	query := fmt.Sprintf(`
	INSERT INTO courses (id, name, description, iduser, creation_date) 
	VALUES ('%s','%s','%s','%s','%s')`, c.ID, c.Name, c.Description, c.UserID, c.CreationDate)

	fmt.Println(query)

	if _, err := cr.store.db.Exec(query); err != nil {
		return err
	}

	return nil
}

// Find ...
func (cr *CourseRepo) Find(id string) (*model.Course, error) {

	c := model.Course{}

	query := fmt.Sprintf(`SELECT * FROM courses WHERE id='%s'`, id)
	if err := cr.store.db.QueryRow(query).Scan(&c.ID, &c.Name, &c.Description, &c.UserID, &c.CreationDate); err != nil {
		return nil, err
	}

	return &c, nil
}

// Read ...
func (cr *CourseRepo) Read(id string) ([]*model.Course, error) {

	query := fmt.Sprintf(`SELECT * FROM courses WHERE iduser='%s'`, id)
	rows, err := cr.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	numOfRows := 0
	query = fmt.Sprintf(`SELECT count(id) FROM courses WHERE iduser='%s'`, id)
	if err := cr.store.db.QueryRow(query).Scan(&numOfRows); err != nil {
		return nil, err
	}

	courses := make([]*model.Course, numOfRows)
	i := 0
	for rows.Next() {
		course := model.Course{}

		err = rows.Scan(&course.ID, &course.Name, &course.Description, &course.UserID, &course.CreationDate)
		if err != nil {
			return nil, err
		}
		courses[i] = &course
		i++
	}
	return courses, nil
}

// Update ...
func (cr *CourseRepo) Update(value string, field string, id string) error {

	query := fmt.Sprintf("UPDATE courses SET `%s`='%s' WHERE id='%s'", field, value, id)
	if _, err := cr.store.db.Exec(query); err != nil {
		return err
	}

	return nil
}

// Delete ...
func (cr *CourseRepo) Delete(id string) error {
	query := fmt.Sprintf("DELETE FROM courses WHERE id='%s'", id)
	if _, err := cr.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}
