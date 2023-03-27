package customsort

import (
	"regexp"
	"sort"
	"strings"

	"github.com/thalesfsp/customerror"
)

//////
// Vars, consts, and types.
//////

const (
	// Asc is the ascending order.
	Asc = "asc"

	// Desc is the descending order.
	Desc = "desc"
)

var (
	// Matches: `key:order,key:order,...`.
	sortPattern = `[a-zA-Z0-9]+:(asc|desc)`

	// Compiled means ultra fast :).
	sortRegex = regexp.MustCompile(sortPattern)
)

// SortMap created to satisfy the Echo's `BindUnmarshaler` interface. Powered by
// the `Sort` type.
type SortMap map[string]string

// UnmarshalParam is the `BindUnmarshaler` implementation.
func (sM *SortMap) UnmarshalParam(src string) error {
	m, err := Sort(src).ToMap()
	if err != nil {
		return err
	}

	*sM = m

	return nil
}

// ToSort converts a SortMap to a raw `Sort` type.
func (sM *SortMap) ToSort() Sort {
	return NewFromMap(*sM)
}

// Sort is the raw sort from the request.
type Sort string

//////
// Helpers.
//////

// match returns the matches of the `sortRegex` in the `s` string. If not found,
// it returns an error.
func match(s string) ([]string, error) {
	matches := sortRegex.FindAllString(s, -1)

	if len(matches) == 0 {
		return nil, customerror.NewInvalidError("sort format. Expected: `key:order,key:order`")
	}

	return matches, nil
}

//////
// Methods.
//////

// String is the Stringer interface implementation.
func (s Sort) String() string {
	return string(s)
}

// IsValid checks if the `s` is formatted correctly.
func (s Sort) IsValid() bool {
	_, err := match(string(s))

	return err == nil
}

// ToAnyString function receives two parameters: `desiredBetweenKVSeparator`
// and `desiredBetweenEntriesSeparator`, which are used to format the `sort`
// output. It uses `sortRegex` to match and validate `fieldName:order` pairs.
//
// The function returns an error if no matches were found, indicating that the
// `sort` string format is not correct. The error message specifies the expected
// format -> `key:order`. If the `sort` string is formatted correctly, it joins
// the pairs using the `desiredBetweenEntriesSeparator`.
//
// Finally, it replaces the `:` separator with the `desiredBetweenKVSeparator`.
// If the function executes successfully, it returns the formatted sort string
// along with a nil error value.
func (s Sort) ToAnyString(desiredBetweenKVSeparator, desiredBetweenEntriesSeparator string) (string, error) {
	matches, err := match(string(s))
	if err != nil {
		return "", err
	}

	entries := strings.Join(matches, desiredBetweenEntriesSeparator)

	return strings.ReplaceAll(entries, ":", desiredBetweenKVSeparator), nil
}

// ToMap converts and validates `s` to a map[fieldName]order.
func (s Sort) ToMap() (map[string]string, error) {
	matches, err := match(string(s))
	if err != nil {
		return nil, customerror.NewFailedToError("match sort pattern", customerror.WithError(err))
	}

	sortMap := make(map[string]string)

	for _, pair := range matches {
		fieldOrder := strings.Split(pair, ":")

		sortMap[fieldOrder[0]] = fieldOrder[1]
	}

	return sortMap, nil
}

//////
// Exported functionalities.
//////

// NewFromMap creates a new `Sort` from a map[fieldName]order. It first sorts
// the map by the keys, then it joins the pairs using the `,` separator.
func NewFromMap(m map[string]string) Sort {
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	pairs := make([]string, 0, len(keys))

	for _, k := range keys {
		pairs = append(pairs, k+":"+m[k])
	}

	return Sort(strings.Join(pairs, ","))
}
