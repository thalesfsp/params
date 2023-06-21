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
	Any interface{} `json:"any" query:"any" param:"any" form:"any" validate:"omitempty,gt=0"`

	// ID of the resource.
	ID string `json:"id" query:"id" param:"id" form:"id" validate:"required"`

	// TTL is the time-to-live.
	//
	// NOTE: Not all storage supports that.
	TTL time.Duration `json:"ttl,omitempty" query:"ttl" param:"ttl" form:"ttl" validate:"omitempty,gt=0"`
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
