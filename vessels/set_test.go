package vessels

import (
	"slices"
	"testing"
)

func TestNewSet(t *testing.T) {
	s := NewSet[int]()
	expect(t, len(s), 0)
	s = NewSet(1, 2, 3)
	expect(t, len(s), 3)
}

func TestSetLen(t *testing.T) {
	s := NewSet(1, 2, 3)
	expect(t, s.Len(), len(s))
}

func TestSetContains(t *testing.T) {
	s := NewSet(1, 2, 3)
	expect(t, s.Contains(2), true)
	expect(t, s.Contains(9), false)
}

func TestSetContainsAll(t *testing.T) {
	s := NewSet(1, 2, 3, 4)
	expect(t, s.ContainsAll(2, 3), true)
	expect(t, s.ContainsAll(4, 5), false)
}

func TestSetContainsAny(t *testing.T) {
	s := NewSet(1, 2, 3, 4, 5)
	expect(t, s.ContainsAny(1), true)
	expect(t, s.ContainsAny(1, 3), true)
	expect(t, s.ContainsAny(4, 9), true)
	expect(t, s.ContainsAny(9), false)
}

func TestSetAdd(t *testing.T) {
	s := NewSet[int]()
	s.Add(9)
	s.Add(8)
	expect(t, s.Len(), 2)
	expect(t, s.ContainsAll(9, 8), true)
}

func TestSetAppend(t *testing.T) {
	s := NewSet[int]()
	s.Append(1, 2, 3)
	expect(t, s.Len(), 3)
	expect(t, s.ContainsAll(1, 2, 3), true)
}

func TestSetDelete(t *testing.T) {
	s := NewSet(1, 2, 3)
	s.Delete(2)
	expect(t, s.ContainsAll(1, 3), true)
	expect(t, s.Contains(2), false)
}

func TestSetClear(t *testing.T) {
	s := NewSet(1, 2, 3)
	s.Clear()
	expect(t, s.Len(), 0)
	expect(t, s.ContainsAny(1, 2, 3), false)
}

func TestSetKeys(t *testing.T) {
	s := NewSet(1, 2, 3)
	k := s.Keys()
	v := s.Values()
	slices.Sort(k)
	slices.Sort(v)
	expect(t, k, []int{1, 2, 3})
	expect(t, v, []int{1, 2, 3})
}

func TestSetEqual(t *testing.T) {
	a := NewSet(1, 2, 3)
	b := NewSet(1, 2, 3)
	c := NewSet(2, 3, 4)
	d := NewSet(2, 3, 4, 5)
	expect(t, a.Equal(b), true)
	expect(t, a.Equal(c), false)
	expect(t, c.Equal(d), false)
}

func TestSetClone(t *testing.T) {
	a := NewSet(1, 2, 3)
	b := a.Clone()
	expect(t, b.ContainsAll(1, 2, 3), true)
}

func TestSetUnion(t *testing.T) {
	a := NewSet(1, 2, 3)
	b := NewSet(4, 5, 6)
	c := SetUnion(a, b)
	expect(t, c.Len(), 6)
	expect(t, c.ContainsAll(1, 2, 3, 4, 5, 6), true)
}

func TestSetIntersection(t *testing.T) {
	a := NewSet(1, 2, 3, 4)
	b := NewSet(2, 3, 4, 5)
	c := SetIntersection(a, b)
	expect(t, c.Len(), 3)
	expect(t, c.ContainsAll(2, 3, 4), true)
	expect(t, c.ContainsAny(1, 5), false)

	a = NewSet(1, 2, 3)
	b = NewSet(2, 3, 4, 5)
	c = SetIntersection(a, b)
	expect(t, c.Len(), 2)
	expect(t, c.ContainsAll(2, 3), true)
	expect(t, c.ContainsAny(1, 4, 5), false)
	c = SetIntersection(b, a)
	expect(t, c.Len(), 2)
	expect(t, c.ContainsAll(2, 3), true)
	expect(t, c.ContainsAny(1, 4, 5), false)
}

func TestSetDifference(t *testing.T) {
	a := NewSet(1, 2, 3, 4, 5)
	b := NewSet(2, 3, 4, 9)
	c := SetDifference(a, b)
	expect(t, c.Len(), 2)
	expect(t, c.ContainsAll(1, 5), true)
}

func TestSetSymmetricDifference(t *testing.T) {
	a := NewSet(1, 2, 3, 4, 5)
	b := NewSet(2, 3, 4, 9)
	c := SetSymmetricDifference(a, b)
	expect(t, c.Len(), 3)
	expect(t, c.ContainsAll(1, 5, 9), true)
}

func TestSetForEach(t *testing.T) {
	a := NewSet(1, 2, 3)
	value := 0
	a.ForEach(func(x int) {
		value += x
	})
	expect(t, value, 6)
}

func TestSetRangeable(t *testing.T) {
	s := NewSet(1, 2, 3)
	var x []int
	for k := range s {
		x = append(x, k)
	}
	slices.Sort(x)
	expect(t, x, []int{1, 2, 3})
}
