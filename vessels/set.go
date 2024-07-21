package vessels

import "maps"

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](keys ...T) Set[T] {
	s := make(Set[T], len(keys))
	for _, k := range keys {
		s[k] = struct{}{}
	}
	return s
}

func (s Set[T]) contains(k T) bool {
	_, ok := s[k]
	return ok
}

func (s Set[T]) missing(k T) bool {
	_, ok := s[k]
	return !ok
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Contains(k T) bool {
	return s.contains(k)
}

func (s Set[T]) ContainsAll(keys ...T) bool {
	for _, k := range keys {
		if s.missing(k) {
			return false
		}
	}
	return true
}

func (s Set[T]) ContainsAny(keys ...T) bool {
	for _, k := range keys {
		if s.contains(k) {
			return true
		}
	}
	return false
}

func (s Set[T]) Add(k T) {
	s[k] = struct{}{}
}

func (s Set[T]) Append(keys ...T) {
	for _, k := range keys {
		s[k] = struct{}{}
	}
}

func (s Set[T]) Delete(k T) {
	delete(s, k)
}

func (s Set[T]) Clear() {
	clear(s)
}

func (s Set[T]) Keys() []T {
	result := []T{}
	for k := range s {
		result = append(result, k)
	}
	return result
}

func (s Set[T]) Values() []T {
	return s.Keys()
}

func (s Set[T]) Equal(o Set[T]) bool {
	if len(s) != len(o) {
		return false
	}
	for k := range s {
		if o.missing(k) {
			return false
		}
	}
	return true
}

func (s Set[T]) Clone() Set[T] {
	return maps.Clone(s)
}

// a or b or both
func SetUnion[T comparable](a, b Set[T]) Set[T] {
	r := make(map[T]struct{}, len(a)+len(b))
	maps.Copy(r, a)
	maps.Copy(r, b)
	return r
}

// a and b
func SetIntersection[T comparable](a, b Set[T]) Set[T] {
	r := NewSet[T]()
	if b.Len() < a.Len() {
		b, a = a, b // opt: fewer comparisons
	}
	for k := range a {
		if b.contains(k) {
			r[k] = struct{}{}
		}
	}
	return r
}

// from a but not in b
func SetDifference[T comparable](a, b Set[T]) Set[T] {
	r := NewSet[T]()
	for k := range a {
		if b.missing(k) {
			r[k] = struct{}{}
		}
	}
	return r
}

// from a or b but not both
func SetSymmetricDifference[T comparable](a, b Set[T]) Set[T] {
	r := NewSet[T]()
	for k := range a {
		if b.missing(k) {
			r[k] = struct{}{}
		}
	}
	for k := range b {
		if a.missing(k) {
			r[k] = struct{}{}
		}
	}
	return r
}

func (s Set[T]) ForEach(f func(T)) {
	for k := range s {
		f(k)
	}
}
