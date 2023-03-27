package param

import (
	"time"
)

// Update params.
//
// SEE: https://echo.labstack.com/guide/binding/#data-sources on data binding.
type Update struct {
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
	return Validate(u)
}
