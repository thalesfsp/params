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
	// Any use this for cases where you need to pass something down to the
	// adapter.
	Any interface{} `form:"any" json:"any" param:"any" query:"any" validate:"omitempty,gt=0"`

	// Count is used for response, and it is the total number of items.
	Count int64 `default:"0" json:"count" validate:"omitempty,gte=0"`

	// Fields to be included in the response.
	Fields field.Fields `form:"fields" json:"fields" param:"fields" query:"fields" validate:"omitempty,gt=0"`

	// Limit is the maximum number of items to return.
	//
	// NOTE: Default value is 10. It's bad practice to return all items.
	Limit int `default:"10" form:"limit" json:"limit" param:"limit" query:"limit" validate:"omitempty,gt=0"`

	// Offset is the number of items to skip before starting to collect the
	// result set.
	Offset int `default:"0" form:"offset" json:"offset" param:"offset" query:"offset" validate:"omitempty,gte=0"`

	// Routing informs the storage adapter to use a specific routes to access the
	// data.
	Routing []string `form:"routing" json:"routing" param:"routing" query:"routing" validate:"omitempty,gt=0"`

	// Search is the search query.
	Search string `form:"search" json:"search" param:"search" query:"search" validate:"omitempty,gt=0"`

	// SearchRaw if set is passed right way to the adapter.
	//
	// NOTE: NOT all adapters support this feature.
	SearchRaw string `form:"search_raw" json:"search_raw" param:"search_raw" query:"search_raw" validate:"omitempty,gt=0"`

	// Sort is the sort query.
	Sort customsort.SortMap `form:"sort" json:"sort" param:"sort" query:"sort" validate:"omitempty,gt=0"`
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
