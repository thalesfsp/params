package param

// Get params.
//
// SEE: https://echo.labstack.com/guide/binding/#data-sources on data binding.
type Get struct {
	// ID of the resource to delete.
	ID string `json:"id" query:"id" param:"id" form:"id" validate:"required"`

	// Fields to be included in the response.
	Fields Fields `json:"fields" query:"fields" param:"fields" form:"fields" validate:"omitempty,gt=0"`
}

// Process the `default` -> `env` -> `validate` struct's fields tags.
//
// NOTE: This function just for testing purpose.
func (g *Get) Process() error {
	return Validate(g)
}
