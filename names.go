package names

import (
	"math"
	"strings"
	"unicode/utf8"

	"github.com/Zyl9393/strut"
)

// Names is a set of strings which allows for substring search using Find().
type Names struct {
	knownNames                           map[string]map[string]int
	shortQueryToPotentiallyMatchingNames map[string][]string
	maxLookupSubstringRuneCount          int
	isRemoveSupported                    bool
}

// Creates and returns a new *Names which maintains lookup lists for all substrings up to a length of maxLookupSubstringRuneCount runes.
// Increasing maxLookupSubstringRuneCount will increase memory usage. However, it will also increase performance of Find()
// as well as make overall memory usage more predictable. If you are unsure, start with a value of 1, which is the minimum.
// Set supportRemove to true if you need to call Remove().
func New(maxLookupSubstringRuneCount int, supportRemove bool) *Names {
	if maxLookupSubstringRuneCount < 1 {
		panic("maxLookupSubstringRuneCount must be greater than zero")
	}
	return &Names{
		knownNames:                           make(map[string]map[string]int),
		shortQueryToPotentiallyMatchingNames: make(map[string][]string),
		maxLookupSubstringRuneCount:          maxLookupSubstringRuneCount,
		isRemoveSupported:                    supportRemove,
	}
}

// Returns the maxLookupSubstringRuneCount which was passed to New().
func (names *Names) MaxLookupSubstringRuneCount() int {
	return names.maxLookupSubstringRuneCount
}

// IsRemoveSupported returns true if names can handle calls of Remove() without panicking.
func (names *Names) IsRemoveSupported() bool {
	return names.isRemoveSupported
}

// Remove removes name from names in O(n) with n = len(name)*maxLookupSubstringRuneCount.
// If names was constructed by passing false for supportRemove in a call to New(), this method panics.
func (names *Names) Remove(name string) bool {
	if _, ok := names.knownNames[name]; !ok {
		return false
	}
	if !names.isRemoveSupported {
		panic("supportRemove needs to be set to true in call to New() when Remove() needs to be called")
	}

	maxRuneCount := names.maxLookupSubstringRuneCount
	runeCountName := utf8.RuneCountInString(name)
	if runeCountName < maxRuneCount {
		maxRuneCount = runeCountName
	}

	for k, v := range names.knownNames[name] {
		candidates := names.shortQueryToPotentiallyMatchingNames[k]
		candidates[v] = candidates[len(candidates)-1]
		names.shortQueryToPotentiallyMatchingNames[k] = candidates[:len(candidates)-1]
		if len(candidates) == 0 {
			delete(names.shortQueryToPotentiallyMatchingNames, k)
		} else {
			names.knownNames[candidates[v]][k] = v
		}
	}

	delete(names.knownNames, name)

	return true
}

// Returns the amount of strings stored in names.
func (names *Names) Size() int {
	return len(names.knownNames)
}

// Add adds name to names and returns true if name was not already contained in names; otherwise, Add does nothing and returns false.
// Add performs in amortized O(n) with n = len(name)*maxLookupSubstringRuneCount.
func (names *Names) Add(name string) bool {
	if names.Contains(name) {
		return false
	}

	maxRuneCount := names.maxLookupSubstringRuneCount
	runeCountName := utf8.RuneCountInString(name)
	if runeCountName < maxRuneCount {
		maxRuneCount = runeCountName
	}

	if names.isRemoveSupported {
		capacity := 0
		strut.IterateSubstringsUnique(name, 1, maxRuneCount, func(from, to int) {
			capacity++
		})
		names.knownNames[name] = make(map[string]int, capacity)
	} else {
		names.knownNames[name] = nil
	}

	strut.IterateSubstringsUnique(name, 1, maxRuneCount, func(from, to int) {
		names.add(name[from:to], name)
	})

	return true
}

func (names *Names) add(q string, name string) {
	if names.shortQueryToPotentiallyMatchingNames[q] == nil {
		var capacity int
		switch len(q) {
		case 0:
			panic("len(q) = 0")
		case 1:
			capacity = 15
		case 2:
			capacity = 7
		default:
			capacity = 1
		}
		names.shortQueryToPotentiallyMatchingNames[q] = make([]string, 0, capacity)
	}
	names.shortQueryToPotentiallyMatchingNames[q] = append(names.shortQueryToPotentiallyMatchingNames[q], name)
	if names.isRemoveSupported {
		names.knownNames[name][q] = len(names.shortQueryToPotentiallyMatchingNames[q]) - 1
	}
}

// Returns true if names contains name, false otherwise.
//
// This method executes in constant time, i.e. O(1).
func (names *Names) Contains(name string) bool {
	_, ok := names.knownNames[name]
	return ok
}

// Returns a slice of all strings containing query. If buffer is non-nil, Find will use it to its full capacity
// for writing the result to, in an attempt to prevent unnecessary memory allocations.
//
// When disregarding writes to the result slice, Find performs in O(1) when utf8.RuneCountInString(query) <= maxLookupSubstringRuneCount.
// Otherwise, Find performs in O(n) with n = length of query + total length of all strings in names containing the least common substring
// of query which is no longer than maxLookupSubstringRuneCount.
func (names *Names) Find(query string, buffer []string) []string {
	if buffer != nil {
		buffer = buffer[:0]
	}
	queryRuneCount := utf8.RuneCountInString(query)
	if query == "" {
		if cap(buffer) < len(names.knownNames) {
			buffer = make([]string, 0, len(names.knownNames))
		}
		for k := range names.knownNames {
			buffer = append(buffer, k)
		}
	} else if queryRuneCount <= names.maxLookupSubstringRuneCount {
		result := names.shortQueryToPotentiallyMatchingNames[query]
		if cap(buffer) < len(result) {
			buffer = make([]string, len(result))
		} else {
			buffer = buffer[:len(result)]
		}
		copy(buffer, result)
	} else {
		searchSlice := names.getSearchSlice(query, queryRuneCount)
		if searchSlice != nil {
			for _, name := range searchSlice {
				if strings.Contains(name, query) {
					buffer = append(buffer, name)
				}
			}
		}
	}
	if buffer == nil {
		return make([]string, 0)
	}
	return buffer
}

func (names *Names) getSearchSlice(query string, queryRuneCount int) (searchSlice []string) {
	best := math.MaxInt
	indices := make([]int, queryRuneCount+1)
	runeIndex := 0
	var i int
	for i = range query + "." {
		indices[runeIndex] = i
		if runeIndex >= names.maxLookupSubstringRuneCount {
			candidate := names.shortQueryToPotentiallyMatchingNames[query[indices[runeIndex-names.maxLookupSubstringRuneCount]:i]]
			if len(candidate) <= best {
				searchSlice = candidate
				best = len(candidate)
			}
		}
		runeIndex++
	}
	return searchSlice
}

// Returns the amount of strings in names which would effectively get scanned to find matches when calling Find().
func (names *Names) NumSearchNames(query string) int {
	queryRuneCount := utf8.RuneCountInString(query)
	if queryRuneCount <= names.maxLookupSubstringRuneCount {
		return 0
	} else {
		return len(names.getSearchSlice(query, queryRuneCount))
	}
}
