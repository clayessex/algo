package vessels

import (
	"testing"

	"github.com/clayessex/algo/expected"
)

func TestNewOrderedMap(t *testing.T) {
	om := NewOrderedMap[int, int]()
	expect(t, om.Len(), 0)
}

func TestLen(t *testing.T) {
	om := NewOrderedMap[int, int]()
	expect(t, om.Len(), 0)
}

func TestInsert(t *testing.T) {
	m := NewOrderedMap[string, int]()
	m.Insert("a", 9)
	m.Insert("b", 8)
	m.Insert("c", 7)
	expect(t, m.Len(), 3)
	expect(t, m.First(), "c") // fifo
	expect(t, m.Last(), "a")
}

func TestContains(t *testing.T) {
	m := NewOrderedMap[int, int]()
	m.Insert(1, 9)
	m.Insert(2, 8)
	expect(t, m.Contains(3), false)
	expect(t, m.Contains(1), true)
	expect(t, m.Contains(2), true)
}

func TestValue(tt *testing.T) {
	m := NewOrderedMap[int, int]()
	m.Insert(1, 9)
	m.Insert(2, 8)

	t := expected.New(tt)
	t.ExpectOk(m.Value(2)).ToBe(8)
}
