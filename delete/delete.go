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
	ID string `form:"id" json:"id" param:"id" query:"id" validate:"required"`

	// Routing informs the storage adapter to use a specific routes to access the
	// data.
	Routing []string `form:"routing" json:"routing" param:"routing" query:"routing" validate:"omitempty,gt=0"`

	// Soft delete if specified.
	Soft bool `default:"false" form:"soft" json:"soft" param:"soft" query:"soft"`

	// Target to delete.
	Target string `form:"target" json:"target" param:"target" query:"target"`
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
