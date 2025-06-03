package customsort

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
