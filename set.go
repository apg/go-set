package set

type Set struct {
	elements map[interface{}]struct{}
}

// New creates and returns a new, empty set.
func New(xs ...interface{}) *Set {
	n := make(map[interface{}]struct{})
	for _, x := range xs {
		n[x] = struct{}{}
	}
	return &Set{n}
}

// Copy returns a distinct copy of `s`.
func (s *Set) Copy() *Set {
	x := New()
	for k, _ := range s.elements {
		x.elements[k] = struct{}{}
	}

	return x
}

// Add adds the element `x` to the set `s` and returns true if `x` was added.
func (s *Set) Add(x interface{}) bool {
	if _, exists := s.elements[x]; !exists {
		s.elements[x] = struct{}{}
		return true
	}
	return false
}

// Delete removes the element `x` from the set `s` and returns true if `x` was found.
func (s *Set) Delete(x interface{}) bool {
	if _, exists := s.elements[x]; exists {
		delete(s.elements, x)
		return true
	}
	return false
}

// Member returns true if `x` is contained in `s`.
func (s *Set) Member(x interface{}) bool {
	_, exists := s.elements[x]
	return exists
}

// Merge adds the elements in `x` to `s`.
func (s *Set) Merge(x *Set) {
	for k, _ := range x.elements {
		s.elements[k] = struct{}{}
	}
}

// Discard removes the elements of `x` from `s`.
func (s *Set) Discard(x *Set) {
	for k, _ := range x.elements {
		delete(s.elements, k)
	}
}

// Difference returns a new set containing the difference between `s` and `x`.
func (s *Set) Difference(x *Set) *Set {
	n := s.Copy()
	for k, _ := range x.elements {
		delete(n.elements, k)
	}

	return n
}

// Union returns a new set composed of the union of `s` and `x`.
func (s *Set) Union(x *Set) *Set {
	n := New()
	for k, _ := range s.elements {
		n.elements[k] = struct{}{}
	}

	for k, _ := range x.elements {
		n.elements[k] = struct{}{}
	}

	return n
}

// Intersection returns a new set composed of the intersection between `s` and `x`.
func (s *Set) Intersection(x *Set) *Set {
	n := New()
	for k1, _ := range x.elements {
		for k2, _ := range s.elements {
			if k1 == k2 {
				n.elements[k1] = struct{}{}
			}
		}
	}

	return n
}

// Superset returns true if `s` is a superset of `x`.
func (s *Set) Superset(x *Set) bool {
	return x.Subset(s)
}

// Subset returns true of `s` is a subset of `x`.
func (s *Set) Subset(x *Set) bool {
	for k, _ := range s.elements {
		if _, exists := x.elements[k]; !exists {
			return false
		}
	}

	return true
}

// Len returns the cardinality of the set.
func (s *Set) Len() int {
	return len(s.elements)
}

// Empty returns true if set is empty
func (s *Set) Empty() bool {
	return len(s.elements) == 0
}

// Freeze returns a slice representing the underlying set data.
func (s *Set) Freeze() []interface{} {
	x := make([]interface{}, len(s.elements))
	i := 0
	for k, _ := range s.elements {
		x[i] = k
		i++
	}

	return x
}
