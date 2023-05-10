package list

import (
	"github.com/thalesfsp/params/customsort"
	"github.com/thalesfsp/params/field"
	"github.com/thalesfsp/params/internal/shared"
	"github.com/thalesfsp/validation"
)

// List params.
//
// SEE: https://echo.labstack.com/guide/binding/#data-sources on data binding.
type List struct {
	// Count is used for response, and it is the total number of items.
	Count int64 `json:"count" validate:"omitempty,gte=0" default:"0"`

	// Fields to be included in the response.
	Fields field.Fields `json:"fields" query:"fields" param:"fields" form:"fields" validate:"omitempty,gt=0"`

	// Limit is the maximum number of items to return.
	//
	// NOTE: Default value is 10. It's bad practice to return all items.
	Limit int `json:"limit" query:"limit" param:"limit" form:"limit" validate:"omitempty,gt=0" default:"10"`

	// Offset is the number of items to skip before starting to collect the
	// result set.
	Offset int `json:"offset" query:"offset" param:"offset" form:"offset" validate:"omitempty,gte=0" default:"0"`

	// Search is the search query.
	Search string `json:"search" query:"search" param:"search" form:"search" validate:"omitempty,gt=0"`

	// SearchRaw if set is passed right way to the adapter.
	//
	// NOTE: NOT all adapters support this feature.
	SearchRaw string `json:"search_raw" query:"search_raw" param:"search_raw" form:"search_raw" validate:"omitempty,gt=0"`

	// Sort is the sort query.
	Sort customsort.SortMap `json:"sort" query:"sort" param:"sort" form:"sort" validate:"omitempty,gt=0"`
}

// Process the `default` -> `env` -> `validate` struct's fields tags.
//
// NOTE: This function just for testing purpose.
func (l *List) Process() error {
	return validation.Validate(l)
}

// New creates a new param.
func New() (*List, error) {
	c := &List{}

	if err := shared.Process(c); err != nil {
		return nil, err
	}

	return c, nil
}
