package param

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
	return Validate(c)
}
