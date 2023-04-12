package retrieve

import (
	"github.com/thalesfsp/params/field"
	"github.com/thalesfsp/params/internal/shared"
	"github.com/thalesfsp/validation"
)

// Retrieve params.
//
// SEE: https://echo.labstack.com/guide/binding/#data-sources on data binding.
type Retrieve struct {
	// ID of the resource to delete.
	ID string `json:"id" query:"id" param:"id" form:"id" validate:"required"`

	// Fields to be included in the response.
	Fields field.Fields `json:"fields" query:"fields" param:"fields" form:"fields" validate:"omitempty,gt=0"`
}

// Process the `default` -> `env` -> `validate` struct's fields tags.
//
// NOTE: This function just for testing purpose.
func (g *Retrieve) Process() error {
	return validation.Validate(g)
}

// New creates a new param.
func New() (*Retrieve, error) {
	c := &Retrieve{}

	if err := shared.Process(c); err != nil {
		return nil, err
	}

	return c, nil
}
