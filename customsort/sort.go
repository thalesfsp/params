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
	sortPattern = `[a-zA-Z0-9.-]+:(asc|desc)`

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
func match(r *regexp.Regexp, s string) ([]string, error) {
	matches := r.FindAllString(s, -1)

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
	_, err := match(sortRegex, string(s))

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
	// First normalize the input by replacing hyphens with commas
	normalizedStr := strings.ReplaceAll(string(s), "-", ",")

	matches, err := match(sortRegex, normalizedStr)
	if err != nil {
		return "", err
	}

	formattedMatches := make([]string, len(matches))
	for i, match := range matches {
		formattedMatches[i] = strings.ReplaceAll(match, ":", desiredBetweenKVSeparator)
	}

	return strings.Join(formattedMatches, desiredBetweenEntriesSeparator), nil
}

// ToMap converts and validates `s` to a map[fieldName]order.
func (s Sort) ToMap() (map[string]string, error) {
	matches, err := match(sortRegex, string(s))
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

// NewFromString creates a new Sort instance from a string representation of sort criteria.
// It supports flexible sort direction symbols:
//   - Both ascSymbol and descSymbol can be defined: "+name,-age"
//   - Only ascSymbol defined: "+name,age" (where age defaults to desc)
//   - Only descSymbol defined: "name,-age" (where name defaults to asc)
//   - Neither can be undefined (will return error)
//
// Parameters:
//   - s: The input string containing sort criteria
//   - betweenEntriesSeparator: Character used to separate multiple sort fields
//   - ascSymbol: Optional symbol for ascending sort (e.g., "+"). If empty and descSymbol
//     is defined, ascending is implicit (no prefix)
//   - descSymbol: Optional symbol for descending sort (e.g., "-"). If empty and ascSymbol
//     is defined, descending is implicit (no prefix)
//
// Returns:
//   - Sort: A properly formatted Sort instance ("field:asc,field:desc")
//   - error: Invalid input errors for empty string, empty entries, invalid parts, or when
//     both symbols are undefined
//
// Examples:
//
//	NewFromString("name:asc,age:desc", ",", "asc", "desc") // name:asc,age:desc
//	NewFromString("name,-age", ",", "", "-")               // name:asc,age:desc
//	NewFromString("+name,age", ",", "+", "")               // name:asc,age:desc
//	NewFromString("+name,-age", ",", "+", "-")             // name:asc,age:desc
//
//nolint:gocognit,gocritic
func NewFromString(
	s,
	betweenEntriesSeparator string,
	ascSymbol, descSymbol string,
) (Sort, error) {
	// Validate that at least one symbol is defined
	if ascSymbol == "" && descSymbol == "" {
		return "", customerror.NewInvalidError("at least one of ascSymbol or descSymbol must be defined")
	}

	// Validate the input string
	if s == "" {
		return "", customerror.NewInvalidError("sort string")
	}

	// Break `s` using the `betweenEntriesSeparator`
	entries := strings.Split(s, betweenEntriesSeparator)

	// Create an accumulator to store the formatted entries
	accumulator := make([]string, 0, len(entries))

	// Iterate over the entries
	for _, entry := range entries {
		// Trim any whitespace
		entry = strings.TrimSpace(entry)

		// Validate the entry
		if entry == "" {
			return "", customerror.NewInvalidError("sort entry")
		}

		// Case 1: Already in the valid `sortRegex` format.
		if sortRegex.MatchString(entry) {
			accumulator = append(accumulator, entry)

			continue
		}

		// Case 2: Both symbols defined
		if ascSymbol != "" && descSymbol != "" {
			if strings.HasPrefix(entry, ascSymbol) {
				entry = strings.TrimPrefix(entry, ascSymbol)
				if entry == "" {
					return "", customerror.NewInvalidError("sort field cannot be empty")
				}

				accumulator = append(accumulator, entry+":"+Asc)
			} else if strings.HasPrefix(entry, descSymbol) {
				entry = strings.TrimPrefix(entry, descSymbol)
				if entry == "" {
					return "", customerror.NewInvalidError("sort field cannot be empty")
				}

				accumulator = append(accumulator, entry+":"+Desc)
			} else {
				return "", customerror.NewInvalidError("sort direction symbol required when both symbols are defined")
			}

			continue
		}

		// Case 3: Only ascSymbol defined
		if ascSymbol != "" {
			if strings.HasPrefix(entry, ascSymbol) {
				entry = strings.TrimPrefix(entry, ascSymbol)
				if entry == "" {
					return "", customerror.NewInvalidError("sort field cannot be empty")
				}

				accumulator = append(accumulator, entry+":"+Asc)
			} else {
				accumulator = append(accumulator, entry+":"+Desc)
			}

			continue
		}

		// Case 4: Only descSymbol defined
		if strings.HasPrefix(entry, descSymbol) {
			entry = strings.TrimPrefix(entry, descSymbol)
			if entry == "" {
				return "", customerror.NewInvalidError("sort field cannot be empty")
			}

			accumulator = append(accumulator, entry+":"+Desc)
		} else {
			accumulator = append(accumulator, entry+":"+Asc)
		}
	}

	// Join the accumulator using the `,` separator
	return Sort(strings.Join(accumulator, ",")), nil
}
