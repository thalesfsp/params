package param

// ListResult response params.
type ListResult[T any] struct {
	Meta    List `json:"meta" validate:"required,dive,required"`
	Results []T  `json:"results"`
}

// Process the `default` -> `env` -> `validate` struct's fields tags.
//
// NOTE: This function just for testing purpose.
func (p *ListResult[T]) Process() error {
	return Validate(p)
}
