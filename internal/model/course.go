package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Course ...
type Course struct {
	ID           string    `json:"id" validate:"required"`
	Name         string    `json:"name" validate:"required"`
	Description  string    `json:"description"`
	UserID       string    `json:"user_id" validate:"required"`
	CreationDate time.Time `json:"creation_date" validate:"required"`
}

// Validate ...
func (c *Course) Validate() error {

	if err := validator.New().Struct(c); err != nil {
		return err
	}
	return nil
}
