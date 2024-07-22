package vessels

import "maps"

type Set[T comparable] map[T]struct{}

// Create a new Set of type T and (optionally) fill it with the given elements
func NewSet[T comparable](elements ...T) Set[T] {
	s := make(Set[T], len(elements))
	for _, k := range elements {
		s[k] = struct{}{}
	}
	return s
}

// Set contains el
func (s Set[T]) contains(el T) bool {
	_, ok := s[el]
	return ok
}

// Set does not contain el
func (s Set[T]) missing(el T) bool {
	_, ok := s[el]
	return !ok
}

// Number of elements in the Set
func (s Set[T]) Len() int {
	return len(s)
}

// True if the Set contains el
func (s Set[T]) Contains(el T) bool {
	return s.contains(el)
}

// True if the Set contains each of the given elements
func (s Set[T]) ContainsAll(elements ...T) bool {
	for _, el := range elements {
		if s.missing(el) {
			return false
		}
	}
	return true
}

// True if the Set contains any of the given elements
func (s Set[T]) ContainsAny(elements ...T) bool {
	for _, el := range elements {
		if s.contains(el) {
			return true
		}
	}
	return false
}

// Add an element to the Set
func (s Set[T]) Add(el T) {
	s[el] = struct{}{}
}

// Add several elements to the Set
func (s Set[T]) Append(elements ...T) {
	for _, el := range elements {
		s[el] = struct{}{}
	}
}

// Remove an element from the Set
func (s Set[T]) Delete(el T) {
	delete(s, el)
}

// Remove all elements from the Set
func (s Set[T]) Clear() {
	clear(s)
}

// Create a slice of type T containing all of the elements in the Set
func (s Set[T]) Keys() []T {
	result := []T{}
	for el := range s {
		result = append(result, el)
	}
	return result
}

// Create a slice of type T containing all of the elements in the Set
// Alias for Keys()
func (s Set[T]) Values() []T {
	return s.Keys()
}

// Compare the elements in the Set to those of parameter 'o' and return true if
// both Sets are of equal length and contain all of the same elements
func (s Set[T]) Equal(o Set[T]) bool {
	if len(s) != len(o) {
		return false
	}
	for el := range s {
		if o.missing(el) {
			return false
		}
	}
	return true
}

// Return a clone of the Set
func (s Set[T]) Clone() Set[T] {
	return maps.Clone(s)
}

// Return a new Set containing the elements in either Set a or Set b or both
func SetUnion[T comparable](a, b Set[T]) Set[T] {
	r := make(map[T]struct{}, len(a)+len(b))
	maps.Copy(r, a)
	maps.Copy(r, b)
	return r
}

// Return a new Set containing the elements in both Set a and Set b
func SetIntersection[T comparable](a, b Set[T]) Set[T] {
	r := NewSet[T]()
	if b.Len() < a.Len() {
		b, a = a, b // opt: fewer comparisons
	}
	for el := range a {
		if b.contains(el) {
			r[el] = struct{}{}
		}
	}
	return r
}

// Return a new Set containing the elements from Set a that are not in Set b
func SetDifference[T comparable](a, b Set[T]) Set[T] {
	r := NewSet[T]()
	for el := range a {
		if b.missing(el) {
			r[el] = struct{}{}
		}
	}
	return r
}

// Return a new Set containing the elements from Set a or Set b, but not both
func SetSymmetricDifference[T comparable](a, b Set[T]) Set[T] {
	r := NewSet[T]()
	for el := range a {
		if b.missing(el) {
			r[el] = struct{}{}
		}
	}
	for el := range b {
		if a.missing(el) {
			r[el] = struct{}{}
		}
	}
	return r
}

// Run the function f against each element of the Set
func (s Set[T]) ForEach(f func(T)) {
	for el := range s {
		f(el)
	}
}
