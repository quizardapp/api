package sqlstore

import (
	"fmt"

	"github.com/quizardapp/auth-api/internal/model"
)

// CardRepo ...
type CardRepo struct {
	store *SQLStore
}

// Create ...
func (cr *CardRepo) Create(c *model.Card) error {
	if err := c.Validate(); err != nil {
		return err
	}

	query := fmt.Sprintf(`
	INSERT INTO cards (id, name, content, idmodule, creation_date) 
	VALUES ('%s','%s','%s','%s','%s')`, c.ID, c.Name, c.Content, c.ModuleID, c.CreationDate)
	if _, err := cr.store.db.Exec(query); err != nil {
		return err
	}

	return nil
}

// Find ...
func (cr *CardRepo) Find(id string) (*model.Card, error) {

	c := model.Card{}

	query := fmt.Sprintf(`SELECT * FROM modules WHERE id='%s'`, id)
	if err := cr.store.db.QueryRow(query).Scan(&c.ID, &c.Name, &c.Content, &c.ModuleID, &c.CreationDate); err != nil {
		return nil, err
	}

	return &c, nil
}

// Read ...
func (cr *CardRepo) Read(id string) ([]*model.Card, error) {
	query := fmt.Sprintf(`SELECT * FROM cards WHERE idmodule='%s'`, id)
	rows, err := cr.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	numOfRows := 0
	query = fmt.Sprintf(`SELECT count(id) FROM cards WHERE idmodule='%s'`, id)
	if err := cr.store.db.QueryRow(query).Scan(&numOfRows); err != nil {
		return nil, err
	}

	fmt.Println(numOfRows)

	cards := make([]*model.Card, numOfRows)
	i := 0
	for rows.Next() {
		card := model.Card{}

		err = rows.Scan(&card.ID, &card.Name, &card.Content, &card.ModuleID, &card.CreationDate)
		if err != nil {
			return nil, err
		}
		cards[i] = &card
		i++
	}
	return cards, nil
}

// Update ...
func (cr *CardRepo) Update(value string, field string, id string) error {
	query := fmt.Sprintf("UPDATE cards SET `%s`='%s' WHERE id='%s'", field, value, id)
	if _, err := cr.store.db.Exec(query); err != nil {
		return err
	}

	return nil
}

// Delete ...
func (cr *CardRepo) Delete(id string) error {
	query := fmt.Sprintf("DELETE FROM cards WHERE id='%s'", id)
	if _, err := cr.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}
