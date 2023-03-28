package field

import (
	"strings"

	"github.com/thalesfsp/params/internal/shared"
)

//////
// Vars, consts, and types.
//////

// Fields to be included in the response.
type Fields []string

//////
// Methods.
//////

// String is the Stringer interface implementation.
func (f Fields) String() string {
	return strings.Join(f, ",")
}

// UnmarshalParam is the `BindUnmarshaler` implementation.
func (f *Fields) UnmarshalParam(src string) error {
	// Split the `src` string by `,`.
	m := Fields(strings.Split(src, ","))

	*f = m

	return nil
}

// ToElasticSearch converts the `f` to a format that can be used in an
// ElasticSearch query: `["field1","field2"]`.
func (f *Fields) ToElasticSearch() (string, error) {
	jsonBytes, err := shared.Marshal(f)
	if err != nil {
		return "", err
	}

	// Convert the JSON bytes to a string
	jsonString := string(jsonBytes)

	return jsonString, nil
}
