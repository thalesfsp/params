package param

// Delete params.
//
// SEE: https://echo.labstack.com/guide/binding/#data-sources on data binding.
type Delete struct {
	// ID of the resource to delete.
	ID string `json:"id" query:"id" param:"id" form:"id" validate:"required"`

	// Soft delete if specified.
	Soft bool `json:"soft" query:"soft" param:"soft" form:"soft" default:"false"`
}

// Process the `default` -> `env` -> `validate` struct's fields tags.
//
// NOTE: This function just for testing purpose.
func (s *Delete) Process() error {
	return Validate(s)
}
