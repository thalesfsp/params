package update

import (
	"time"

	"github.com/thalesfsp/params/internal/shared"
	"github.com/thalesfsp/validation"
)

// Update params.
//
// SEE: https://echo.labstack.com/guide/binding/#data-sources on data binding.
type Update struct {
	// Any use this for cases where you need to pass something down to the
	// adapter.
	Any interface{} `form:"any" json:"any" param:"any" query:"any" validate:"omitempty,gt=0"`

	// ID of the resource.
	ID string `form:"id" json:"id" param:"id" query:"id" validate:"required"`

	// Routing informs the storage adapter to use a specific routes to access the
	// data.
	Routing []string `form:"routing" json:"routing" param:"routing" query:"routing" validate:"omitempty,gt=0"`

	// TTL is the time-to-live.
	//
	// NOTE: Not all storage supports that.
	TTL time.Duration `form:"ttl" json:"ttl,omitempty" param:"ttl" query:"ttl" validate:"omitempty,gt=0"`
}

// Process the `default` -> `env` -> `validate` struct's fields tags.
//
// NOTE: This function just for testing purpose.
func (u *Update) Process() error {
	return validation.Validate(u)
}

// New creates a new param.
func New() (*Update, error) {
	c := &Update{}

	if err := shared.Process(c); err != nil {
		return nil, err
	}

	return c, nil
}
