package names

import (
	"strings"
)

// Sorts case-insensitively, using case-sensitive sorting only where two strings are the same in their lower case forms.
// Use as an argument to slices.SortFunc().
func SortCIFunc(a, b string) bool {
	lowA := strings.ToLower(a)
	lowB := strings.ToLower(b)
	return lowA < lowB || (lowA == lowB && a < b)
}

// Returns a sort function which sorts s for user experience: strings starting with query come first.
func NewSortUXFunc(query string) func(a, b string) bool {
	return func(a, b string) bool {
		return (strings.HasPrefix(a, query) && strings.HasPrefix(b, query) && a < b) || ((strings.HasPrefix(a, query) || a < b) && !strings.HasPrefix(b, query))
	}
}

// Returns a sort function which sorts s case-insensitively, using case-sensitive sorting only where two strings are the same in their lower case forms,
// for user experience: strings case-insensitively starting with query come first.
func NewSortUXCIFunc(query string) func(a, b string) bool {
	query = strings.ToLower(query)
	return func(a, b string) bool {
		lowA := strings.ToLower(a)
		lowB := strings.ToLower(b)
		return (strings.HasPrefix(lowA, query) && strings.HasPrefix(lowB, query) && (lowA < lowB || (lowA == lowB && a < b))) ||
			((strings.HasPrefix(lowA, query) || (lowA < lowB || (lowA == lowB && a < b))) && !strings.HasPrefix(lowB, query))
	}
}

// TODO: Sorting functions which perform natural ordering, i.e. treat "42" as being between "041" and "043" instead of after both.
