package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Card ...
type Card struct {
	ID           string    `json:"id" validate:"required"`
	Name         string    `json:"name" validate:"required"`
	Content      string    `json:"content" validate:"required"`
	UserID       string    `json:"user_id" validate:"required"`
	ModuleID     string    `json:"module_id" validate:"required"`
	CreationDate time.Time `json:"creation_date" validate:"required"`
}

// Validate ...
func (c *Card) Validate() error {

	if err := validator.New().Struct(c); err != nil {
		return err
	}
	return nil
}
