package names

import (
	"reflect"
	"testing"

	"golang.org/x/exp/slices"
)

func TestSortCI(t *testing.T) {
	tests := []struct {
		s        []string
		expected []string
	}{
		{[]string{"a", "A"}, []string{"A", "a"}},
		{[]string{"aa", "a"}, []string{"a", "aa"}},
		{[]string{"a", "B", "c"}, []string{"a", "B", "c"}},
		{[]string{"a", "c", "B"}, []string{"a", "B", "c"}},
		{[]string{"B", "c", "a"}, []string{"a", "B", "c"}},
		{[]string{"c", "B", "a"}, []string{"a", "B", "c"}},
		{[]string{"c", "a", "B"}, []string{"a", "B", "c"}},
		{[]string{"B", "a", "c"}, []string{"a", "B", "c"}},
		{[]string{"B", "B", "a", "c"}, []string{"a", "B", "B", "c"}},
		{[]string{"B", "B", "c", "a"}, []string{"a", "B", "B", "c"}},
		{[]string{"B", "a", "B", "c"}, []string{"a", "B", "B", "c"}},
		{[]string{"B", "c", "B", "a"}, []string{"a", "B", "B", "c"}},
		{[]string{"B", "a", "c", "B"}, []string{"a", "B", "B", "c"}},
		{[]string{"B", "c", "a", "B"}, []string{"a", "B", "B", "c"}},
		{[]string{"a", "B", "c", "B"}, []string{"a", "B", "B", "c"}},
		{[]string{"c", "B", "a", "B"}, []string{"a", "B", "B", "c"}},
		{[]string{"a", "c", "B", "B"}, []string{"a", "B", "B", "c"}},
		{[]string{"c", "a", "B", "B"}, []string{"a", "B", "B", "c"}},
		{[]string{"a", "B", "B", "c"}, []string{"a", "B", "B", "c"}},
		{[]string{"c", "B", "B", "a"}, []string{"a", "B", "B", "c"}},
	}
	for i, test := range tests {
		cp := make([]string, len(test.s))
		copy(cp, test.s)
		slices.SortFunc(cp, SortCIFunc)
		if !reflect.DeepEqual(cp, test.expected) {
			t.Errorf("Test %d: SortCI(%v) resulted in %v. Expected %v.", i, test.s, cp, test.expected)
		}
	}
}

func TestSortUX(t *testing.T) {
	tests := []struct {
		query    string
		s        []string
		expected []string
	}{
		{"k", []string{"a", "A"}, []string{"A", "a"}},
		{"a", []string{"a", "A"}, []string{"a", "A"}},
		{"A", []string{"a", "A"}, []string{"A", "a"}},
		{"k", []string{"a", "B", "c"}, []string{"B", "a", "c"}},
		{"k", []string{"a", "c", "B"}, []string{"B", "a", "c"}},
		{"k", []string{"B", "c", "a"}, []string{"B", "a", "c"}},
		{"k", []string{"c", "B", "a"}, []string{"B", "a", "c"}},
		{"k", []string{"c", "a", "B"}, []string{"B", "a", "c"}},
		{"k", []string{"B", "a", "c"}, []string{"B", "a", "c"}},
		{"a", []string{"a", "B", "c"}, []string{"a", "B", "c"}},
		{"a", []string{"a", "c", "B"}, []string{"a", "B", "c"}},
		{"a", []string{"B", "c", "a"}, []string{"a", "B", "c"}},
		{"a", []string{"c", "B", "a"}, []string{"a", "B", "c"}},
		{"a", []string{"c", "a", "B"}, []string{"a", "B", "c"}},
		{"a", []string{"B", "a", "c"}, []string{"a", "B", "c"}},
		{"A", []string{"a", "B", "c"}, []string{"B", "a", "c"}},
		{"A", []string{"a", "c", "B"}, []string{"B", "a", "c"}},
		{"A", []string{"B", "c", "a"}, []string{"B", "a", "c"}},
		{"A", []string{"c", "B", "a"}, []string{"B", "a", "c"}},
		{"A", []string{"c", "a", "B"}, []string{"B", "a", "c"}},
		{"A", []string{"B", "a", "c"}, []string{"B", "a", "c"}},
		{"k", []string{"B", "B", "a", "c"}, []string{"B", "B", "a", "c"}},
		{"k", []string{"B", "B", "c", "a"}, []string{"B", "B", "a", "c"}},
		{"k", []string{"B", "a", "B", "c"}, []string{"B", "B", "a", "c"}},
		{"k", []string{"B", "c", "B", "a"}, []string{"B", "B", "a", "c"}},
		{"k", []string{"B", "a", "c", "B"}, []string{"B", "B", "a", "c"}},
		{"k", []string{"B", "c", "a", "B"}, []string{"B", "B", "a", "c"}},
		{"k", []string{"a", "B", "c", "B"}, []string{"B", "B", "a", "c"}},
		{"k", []string{"c", "B", "a", "B"}, []string{"B", "B", "a", "c"}},
		{"k", []string{"a", "c", "B", "B"}, []string{"B", "B", "a", "c"}},
		{"k", []string{"c", "a", "B", "B"}, []string{"B", "B", "a", "c"}},
		{"k", []string{"a", "B", "B", "c"}, []string{"B", "B", "a", "c"}},
		{"k", []string{"c", "B", "B", "a"}, []string{"B", "B", "a", "c"}},
		{"a", []string{"B", "B", "a", "c"}, []string{"a", "B", "B", "c"}},
		{"a", []string{"B", "B", "c", "a"}, []string{"a", "B", "B", "c"}},
		{"a", []string{"B", "a", "B", "c"}, []string{"a", "B", "B", "c"}},
		{"a", []string{"B", "c", "B", "a"}, []string{"a", "B", "B", "c"}},
		{"a", []string{"B", "a", "c", "B"}, []string{"a", "B", "B", "c"}},
		{"a", []string{"B", "c", "a", "B"}, []string{"a", "B", "B", "c"}},
		{"a", []string{"a", "B", "c", "B"}, []string{"a", "B", "B", "c"}},
		{"a", []string{"c", "B", "a", "B"}, []string{"a", "B", "B", "c"}},
		{"a", []string{"a", "c", "B", "B"}, []string{"a", "B", "B", "c"}},
		{"a", []string{"c", "a", "B", "B"}, []string{"a", "B", "B", "c"}},
		{"a", []string{"a", "B", "B", "c"}, []string{"a", "B", "B", "c"}},
		{"a", []string{"c", "B", "B", "a"}, []string{"a", "B", "B", "c"}},
		{"A", []string{"B", "B", "a", "c"}, []string{"B", "B", "a", "c"}},
		{"A", []string{"B", "B", "c", "a"}, []string{"B", "B", "a", "c"}},
		{"A", []string{"B", "a", "B", "c"}, []string{"B", "B", "a", "c"}},
		{"A", []string{"B", "c", "B", "a"}, []string{"B", "B", "a", "c"}},
		{"A", []string{"B", "a", "c", "B"}, []string{"B", "B", "a", "c"}},
		{"A", []string{"B", "c", "a", "B"}, []string{"B", "B", "a", "c"}},
		{"A", []string{"a", "B", "c", "B"}, []string{"B", "B", "a", "c"}},
		{"A", []string{"c", "B", "a", "B"}, []string{"B", "B", "a", "c"}},
		{"A", []string{"a", "c", "B", "B"}, []string{"B", "B", "a", "c"}},
		{"A", []string{"c", "a", "B", "B"}, []string{"B", "B", "a", "c"}},
		{"A", []string{"a", "B", "B", "c"}, []string{"B", "B", "a", "c"}},
		{"A", []string{"c", "B", "B", "a"}, []string{"B", "B", "a", "c"}},
	}
	for i, test := range tests {
		cp := make([]string, len(test.s))
		copy(cp, test.s)
		slices.SortFunc(cp, NewSortUXFunc(test.query))
		if !reflect.DeepEqual(cp, test.expected) {
			t.Errorf(`Test %d: SortUX(%v, "%s") resulted in %v. Expected %v.`, i, test.s, test.query, cp, test.expected)
		}
	}
}

func TestSortUXCI(t *testing.T) {
	tests := []struct {
		query    string
		s        []string
		expected []string
	}{
		{"k", []string{"a", "A"}, []string{"A", "a"}},
		{"a", []string{"a", "A"}, []string{"A", "a"}},
		{"A", []string{"a", "A"}, []string{"A", "a"}},
		{"k", []string{"a", "B", "c"}, []string{"a", "B", "c"}},
		{"k", []string{"a", "c", "B"}, []string{"a", "B", "c"}},
		{"k", []string{"B", "c", "a"}, []string{"a", "B", "c"}},
		{"k", []string{"c", "B", "a"}, []string{"a", "B", "c"}},
		{"k", []string{"c", "a", "B"}, []string{"a", "B", "c"}},
		{"k", []string{"B", "a", "c"}, []string{"a", "B", "c"}},
		{"a", []string{"a", "B", "c"}, []string{"a", "B", "c"}},
		{"a", []string{"a", "c", "B"}, []string{"a", "B", "c"}},
		{"a", []string{"B", "c", "a"}, []string{"a", "B", "c"}},
		{"a", []string{"c", "B", "a"}, []string{"a", "B", "c"}},
		{"a", []string{"c", "a", "B"}, []string{"a", "B", "c"}},
		{"a", []string{"B", "a", "c"}, []string{"a", "B", "c"}},
		{"A", []string{"a", "B", "c"}, []string{"a", "B", "c"}},
		{"A", []string{"a", "c", "B"}, []string{"a", "B", "c"}},
		{"A", []string{"B", "c", "a"}, []string{"a", "B", "c"}},
		{"A", []string{"c", "B", "a"}, []string{"a", "B", "c"}},
		{"A", []string{"c", "a", "B"}, []string{"a", "B", "c"}},
		{"A", []string{"B", "a", "c"}, []string{"a", "B", "c"}},
		{"k", []string{"B", "B", "a", "c"}, []string{"a", "B", "B", "c"}},
		{"k", []string{"B", "B", "c", "a"}, []string{"a", "B", "B", "c"}},
		{"k", []string{"B", "a", "B", "c"}, []string{"a", "B", "B", "c"}},
		{"k", []string{"B", "c", "B", "a"}, []string{"a", "B", "B", "c"}},
		{"k", []string{"B", "a", "c", "B"}, []string{"a", "B", "B", "c"}},
		{"k", []string{"B", "c", "a", "B"}, []string{"a", "B", "B", "c"}},
		{"k", []string{"a", "B", "c", "B"}, []string{"a", "B", "B", "c"}},
		{"k", []string{"c", "B", "a", "B"}, []string{"a", "B", "B", "c"}},
		{"k", []string{"a", "c", "B", "B"}, []string{"a", "B", "B", "c"}},
		{"k", []string{"c", "a", "B", "B"}, []string{"a", "B", "B", "c"}},
		{"k", []string{"a", "B", "B", "c"}, []string{"a", "B", "B", "c"}},
		{"k", []string{"c", "B", "B", "a"}, []string{"a", "B", "B", "c"}},
		{"B", []string{"B", "B", "a", "c"}, []string{"B", "B", "a", "c"}},
		{"B", []string{"B", "B", "c", "a"}, []string{"B", "B", "a", "c"}},
		{"B", []string{"B", "a", "B", "c"}, []string{"B", "B", "a", "c"}},
		{"B", []string{"B", "c", "B", "a"}, []string{"B", "B", "a", "c"}},
		{"B", []string{"B", "a", "c", "B"}, []string{"B", "B", "a", "c"}},
		{"B", []string{"B", "c", "a", "B"}, []string{"B", "B", "a", "c"}},
		{"B", []string{"a", "B", "c", "B"}, []string{"B", "B", "a", "c"}},
		{"B", []string{"c", "B", "a", "B"}, []string{"B", "B", "a", "c"}},
		{"B", []string{"a", "c", "B", "B"}, []string{"B", "B", "a", "c"}},
		{"B", []string{"c", "a", "B", "B"}, []string{"B", "B", "a", "c"}},
		{"B", []string{"a", "B", "B", "c"}, []string{"B", "B", "a", "c"}},
		{"B", []string{"c", "B", "B", "a"}, []string{"B", "B", "a", "c"}},
		{"C", []string{"B", "B", "a", "c"}, []string{"c", "a", "B", "B"}},
		{"C", []string{"B", "B", "c", "a"}, []string{"c", "a", "B", "B"}},
		{"C", []string{"B", "a", "B", "c"}, []string{"c", "a", "B", "B"}},
		{"C", []string{"B", "c", "B", "a"}, []string{"c", "a", "B", "B"}},
		{"C", []string{"B", "a", "c", "B"}, []string{"c", "a", "B", "B"}},
		{"C", []string{"B", "c", "a", "B"}, []string{"c", "a", "B", "B"}},
		{"C", []string{"a", "B", "c", "B"}, []string{"c", "a", "B", "B"}},
		{"C", []string{"c", "B", "a", "B"}, []string{"c", "a", "B", "B"}},
		{"C", []string{"a", "c", "B", "B"}, []string{"c", "a", "B", "B"}},
		{"C", []string{"c", "a", "B", "B"}, []string{"c", "a", "B", "B"}},
		{"C", []string{"a", "B", "B", "c"}, []string{"c", "a", "B", "B"}},
		{"C", []string{"c", "B", "B", "a"}, []string{"c", "a", "B", "B"}},
	}
	for i, test := range tests {
		cp := make([]string, len(test.s))
		copy(cp, test.s)
		slices.SortFunc(cp, NewSortUXCIFunc(test.query))
		if !reflect.DeepEqual(cp, test.expected) {
			t.Errorf(`Test %d: SortUXCI(%v, "%s") resulted in %v. Expected %v.`, i, test.s, test.query, cp, test.expected)
		}
	}
}
