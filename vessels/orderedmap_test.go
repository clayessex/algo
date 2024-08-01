package vessels

import "testing"

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
