package listresult

import (
	"github.com/thalesfsp/params/list"
	"github.com/thalesfsp/validation"
)

// ListResult response params.
type ListResult[T any] struct {
	Meta    list.List `json:"meta" validate:"required"`
	Results []T       `json:"results"`
}

// Process the `default` -> `env` -> `validate` struct's fields tags.
//
// NOTE: This function just for testing purpose.
func (p *ListResult[T]) Process() error {
	return validation.Validate(p)
}
