package directed_test

import "slices"

// node definition for testing
type node struct {
	id string
}

// ShouldBeOrdered checks if the given elements are ordered in the slice
func ShouldBeOrdered[T comparable](s []T, items ...T) bool {
	lastPos := -1
	for _, item := range items {
		pos := slices.Index(s, item)
		if pos == -1 {
			return false
		}
		if pos < lastPos {
			return false
		}
		lastPos = pos
	}
	return true
}
