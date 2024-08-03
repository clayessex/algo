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

	x := expected.New(t)
	x.Expect(m.Len()).ToBe(3)
	x.Expect(m.First()).ToBe("c")
	x.Expect(m.Last()).ToBe("a")

	x.ExpectOk(m.Value("a")).ToBe(9)
	m.Insert("a", 3)
	x.Expect(m.Last()).ToBe("a")
	x.ExpectOk(m.Value("a")).ToBe(3)
}

func TestContains(t *testing.T) {
	m := NewOrderedMap[int, int]()
	m.Insert(1, 9)
	m.Insert(2, 8)
	expect(t, m.Contains(3), false)
	expect(t, m.Contains(1), true)
	expect(t, m.Contains(2), true)
}

func TestValue(t *testing.T) {
	m := NewOrderedMap[int, int]()
	m.Insert(1, 9)
	m.Insert(2, 8)
	x := expected.New(t)
	x.ExpectOk(m.Value(2)).ToBe(8)
	x.ExpectNotOk(m.Value(5))
}
