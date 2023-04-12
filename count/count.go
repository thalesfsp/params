package count

import (
	"github.com/thalesfsp/params/internal/shared"
	"github.com/thalesfsp/validation"
)

// Count params.
//
// SEE: https://echo.labstack.com/guide/binding/#data-sources on data binding.
type Count struct {
	// Search is the search query.
	Search string `json:"search" query:"search" param:"search" form:"search" validate:"required"`
}

// Process the `default` -> `env` -> `validate` struct's fields tags.
//
// NOTE: This function just for testing purpose.
func (c *Count) Process() error {
	return validation.Validate(c)
}

// New creates a new param.
func New() (*Count, error) {
	c := &Count{}

	if err := shared.Process(c); err != nil {
		return nil, err
	}

	return c, nil
}
