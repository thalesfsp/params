package delete

import (
	"github.com/thalesfsp/params/internal/shared"
	"github.com/thalesfsp/validation"
)

// Delete params.
//
// SEE: https://echo.labstack.com/guide/binding/#data-sources on data binding.
type Delete struct {
	// ID of the resource to delete.
	ID string `json:"id" query:"id" param:"id" form:"id" validate:"required"`

	// Soft delete if specified.
	Soft bool `json:"soft" query:"soft" param:"soft" form:"soft" default:"false"`

	// Target to delete.
	Target string `json:"target" query:"target" param:"target" form:"target"`
}

// Process the `default` -> `env` -> `validate` struct's fields tags.
//
// NOTE: This function just for testing purpose.
func (s *Delete) Process() error {
	return validation.Validate(s)
}

// New creates a new param.
func New() (*Delete, error) {
	c := &Delete{}

	if err := shared.Process(c); err != nil {
		return nil, err
	}

	return c, nil
}
