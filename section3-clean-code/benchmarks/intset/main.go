package main

// IntSet implements a set of integers
// (set elements are unique).
type IntSet struct {
	elems map[int]struct{}
}

// MakeIntSet creates an empty set.
func MakeIntSet() IntSet {
	return IntSet{make(map[int]struct{})}
}

// Contains reports whether the element is in the set.
func (s IntSet) Contains(elem int) bool {
	_, ok := s.elems[elem]
	return ok
}

// Add adds the element to the set.
// Returns true if the element was added,
// otherwise false (if the element is already in the set).
func (s IntSet) Add(elem int) bool {
	if _, ok := s.elems[elem]; ok {
		return false
	}
	s.elems[elem] = struct{}{}
	return true
}
