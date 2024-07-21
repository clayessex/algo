package vessels

import "maps"

type emptyValue = struct{}

type Set[T comparable] struct {
	m map[T]emptyValue
}

func NewSet[T comparable](keys ...T) *Set[T] {
	s := Set[T]{}
	s.m = make(map[T]emptyValue)
	for _, k := range keys {
		s.m[k] = emptyValue{}
	}
	return &s
}

func (s *Set[T]) contains(k T) bool {
	_, ok := s.m[k]
	return ok
}

func (s *Set[T]) missing(k T) bool {
	_, ok := s.m[k]
	return !ok
}

func (s *Set[T]) Len() int {
	return len(s.m)
}

func (s *Set[T]) Contains(k T) bool {
	return s.contains(k)
}

func (s *Set[T]) ContainsAll(keys ...T) bool {
	for _, k := range keys {
		if s.missing(k) {
			return false
		}
	}
	return true
}

func (s *Set[T]) ContainsAny(keys ...T) bool {
	for _, k := range keys {
		if s.contains(k) {
			return true
		}
	}
	return false
}

func (s *Set[T]) Add(k T) {
	s.m[k] = emptyValue{}
}

func (s *Set[T]) Append(keys ...T) {
	for _, k := range keys {
		s.m[k] = emptyValue{}
	}
}

func (s *Set[T]) Delete(k T) {
	delete(s.m, k)
}

func (s *Set[T]) Clear() {
	clear(s.m)
}

func (s *Set[T]) Keys() []T {
	result := []T{}
	for k := range s.m {
		result = append(result, k)
	}
	return result
}

func (s *Set[T]) Values() []T {
	return s.Keys()
}

func (s *Set[T]) Equal(o *Set[T]) bool {
	if len(s.m) != len(o.m) {
		return false
	}
	for k := range s.m {
		if o.missing(k) {
			return false
		}
	}
	return true
}

func (s *Set[T]) Clone() *Set[T] {
	r := Set[T]{}
	r.m = maps.Clone(s.m)
	return &r
}

// a or b or both
func SetUnion[T comparable](a, b *Set[T]) *Set[T] {
	r := Set[T]{}
	r.m = make(map[T]emptyValue, len(a.m)+len(b.m))
	maps.Copy(r.m, a.m)
	maps.Copy(r.m, b.m)
	return &r
}

// a and b
func SetIntersection[T comparable](a, b *Set[T]) *Set[T] {
	r := NewSet[T]()
	if b.Len() < a.Len() {
		b, a = a, b // opt: fewer comparisons
	}
	for k := range a.m {
		if b.contains(k) {
			r.m[k] = emptyValue{}
		}
	}
	return r
}

// from a but not in b
func SetDifference[T comparable](a, b *Set[T]) *Set[T] {
	r := NewSet[T]()
	for k := range a.m {
		if b.missing(k) {
			r.m[k] = emptyValue{}
		}
	}
	return r
}

// from a or b but not both
func SetSymmetricDifference[T comparable](a, b *Set[T]) *Set[T] {
	r := NewSet[T]()
	for k := range a.m {
		if b.missing(k) {
			r.m[k] = emptyValue{}
		}
	}
	for k := range b.m {
		if a.missing(k) {
			r.m[k] = emptyValue{}
		}
	}
	return r
}

func (s *Set[T]) ForEach(f func(T)) {
	for k := range s.m {
		f(k)
	}
}
