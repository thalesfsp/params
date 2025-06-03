package customsort

// SortSlice created to satisfy the Echo's `BindUnmarshaler` interface. Powered by
// the `Sort` type.
type SortSlice [][]string

// UnmarshalParam is the `BindUnmarshaler` implementation.
func (sS *SortSlice) UnmarshalParam(src string) error {
	s, err := Sort(src).ToSlice()
	if err != nil {
		return err
	}

	*sS = s

	return nil
}

// ToSort converts a SortSlice to a raw `Sort` type.
func (sS *SortSlice) ToSort() Sort {
	return NewFromSlice(*sS)
}
