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
	x.ExpectOk(m.First()).ToBe("c")
	x.ExpectOk(m.Last()).ToBe("a")

	x.ExpectOk(m.Value("a")).ToBe(9)
	m.Insert("a", 3)
	x.ExpectOk(m.Last()).ToBe("a")
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

func TestDelete(t *testing.T) {
	x := expected.New(t)
	m := NewOrderedMap[int, int]()
	m.Insert(1, 9)
	m.Insert(2, 8)
	x.Expect(m.Delete(2)).ToBe(true)
	x.Expect(m.Len()).ToBe(1)
	x.Expect(m.Delete(42)).ToBe(false)
}

func TestPush(t *testing.T) {
	m := NewOrderedMap[int, int]()
	m.Push(1, 9)
	expected.Expect(t, m.Len(), 1)
}

func TestPop(t *testing.T) {
	x := expected.New(t)
	m := NewOrderedMap[int, int]()
	m.Push(1, 9)
	x.ExpectOk(m.Pop()).ToBe(1)
	x.ExpectNotOk(m.Pop())
}

func TestOMNext(t *testing.T) {
	x := expected.New(t)
	m := NewOrderedMap[int, int]()
	x.ExpectNotOk(m.Next(42))

	m.Push(1, 9)
	x.ExpectNotOk(m.Next(1))
	m.Push(2, 9)
	x.ExpectOk(m.Next(1)).ToBe(2)
}
