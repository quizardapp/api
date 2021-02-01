package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Module ...
type Module struct {
	ID           string `json:"id" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Description  string `json:"description"`
	UserID       string `json:"user_id" validate:"required"`
	CourseID     string `json:"course_id"`
	CreationDate time.Time `json:"creation_date" validate:"required"`
}

// Validate ...
func (m *Module) Validate() error {

	if err := validator.New().Struct(m); err != nil {
		return err
	}
	return nil
}
