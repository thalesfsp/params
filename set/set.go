package set

import (
	"time"

	"github.com/thalesfsp/params/internal/shared"
	"github.com/thalesfsp/params/vali"
)

// Set params.
//
// SEE: https://echo.labstack.com/guide/binding/#data-sources on data binding.
type Set struct {
	// ID of the resource.
	ID string `json:"id,omitempty" query:"id" param:"id" form:"id" validate:"omitempty,gt=0"`

	// TTL is the time-to-live.
	//
	// NOTE: Not all storage supports that.
	TTL time.Duration `json:"ttl,omitempty" query:"ttl" param:"ttl" form:"ttl" validate:"omitempty,gt=0"`
}

// Process the `default` -> `env` -> `validate` struct's fields tags.
//
// NOTE: This function just for testing purpose.
func (s *Set) Process() error {
	return vali.Validate(s)
}

// New creates a new param.
func New() (*Set, error) {
	c := &Set{}

	if err := shared.Process(c); err != nil {
		return nil, err
	}

	return c, nil
}
