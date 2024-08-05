package vessels

import (
	"testing"

	"github.com/clayessex/algo/expected"
)

func TestNewOrderedMap(t *testing.T) {
	om := NewOrderedMap[int, int]()
	expect(t, om.Len(), 0)
	om = NewOrderedMap[int, int](10)
	expect(t, om.Len(), 0)
}

func TestOMLen(t *testing.T) {
	om := NewOrderedMap[int, int]()
	expect(t, om.Len(), 0)
}

func TestOMInsert(t *testing.T) {
	m := NewOrderedMap[string, int]()
	m.Insert("a", 9)
	m.Insert("b", 8)
	m.Insert("c", 7)

	x := expected.New(t)
	x.Expect(m.Len()).ToBe(3)
	x.ExpectOk(m.First()).ToBe("a")
	x.ExpectOk(m.Last()).ToBe("c")

	x.ExpectOk(m.Value("a")).ToBe(9)
	m.Insert("a", 3)
	x.ExpectOk(m.First()).ToBe("a")
	x.ExpectOk(m.Last()).ToBe("c")
	x.ExpectOk(m.Value("a")).ToBe(3)
}

func TestOMContains(t *testing.T) {
	m := NewOrderedMap[int, int]()
	m.Insert(1, 9)
	m.Insert(2, 8)
	expect(t, m.Contains(3), false)
	expect(t, m.Contains(1), true)
	expect(t, m.Contains(2), true)
}

func TestOMValue(t *testing.T) {
	m := NewOrderedMap[int, int]()
	m.Insert(1, 9)
	m.Insert(2, 8)
	x := expected.New(t)
	x.ExpectOk(m.Value(2)).ToBe(8)
	x.ExpectNotOk(m.Value(5))
}

func TestOMDelete(t *testing.T) {
	x := expected.New(t)
	m := NewOrderedMap[int, int]()
	m.Insert(1, 9)
	m.Insert(2, 8)
	x.Expect(m.Delete(2)).ToBe(true)
	x.Expect(m.Len()).ToBe(1)
	x.Expect(m.Delete(42)).ToBe(false)
}

func TestOMPush(t *testing.T) {
	m := NewOrderedMap[int, int]()
	m.Push(1, 9)
	expected.Expect(t, m.Len(), 1)
}

func TestOMPop(t *testing.T) {
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

func TestOMPrev(t *testing.T) {
	x := expected.New(t)
	m := NewOrderedMap[int, int]()
	x.ExpectNotOk(m.Prev(42))

	m.Insert(1, 9)
	x.ExpectNotOk(m.Prev(1))
	m.Insert(2, 8)
	x.ExpectOk(m.Prev(2)).ToBe(1)
}

func TestOMFirst(t *testing.T) {
	x := expected.New(t)
	m := NewOrderedMap[int, int]()
	x.ExpectNotOk(m.First())
	m.Push(1, 9)
	x.ExpectOk(m.First()).ToBe(1)
	m.Push(2, 8)
	x.ExpectOk(m.First()).ToBe(1)
}

func TestOMLast(t *testing.T) {
	x := expected.New(t)
	m := NewOrderedMap[int, int]()
	x.ExpectNotOk(m.Last())
	m.Push(1, 9)
	x.ExpectOk(m.Last()).ToBe(1)
	m.Push(2, 8)
	x.ExpectOk(m.Last()).ToBe(2)
}

func TestOMClear(t *testing.T) {
	x := expected.New(t)
	m := NewOrderedMap[int, int]()
	m.Push(1, 9)
	m.Push(2, 8)
	x.Expect(m.Len()).ToBe(2)
	m.Clear()
	x.Expect(m.Len()).ToBe(0)
	x.ExpectNotOk(m.Value(1))
}

func TestOMKeys(t *testing.T) {
	x := expected.New(t)
	m := NewOrderedMap[int, int]()
	m.Push(1, 9)
	m.Push(2, 8)
	m.Push(3, 7)
	x.Expect(m.Keys()).ToBe([]int{1, 2, 3})
}

func TestOMValues(t *testing.T) {
	x := expected.New(t)
	m := NewOrderedMap[int, int]()
	m.Push(1, 9)
	m.Push(2, 8)
	m.Push(3, 7)
	x.Expect(m.Values()).ToBe([]int{9, 8, 7})
}

func TestOMAt(t *testing.T) {
	x := expected.New(t)
	m := NewOrderedMap[int, int]()
	m.Push(1, 9)
	m.Push(2, 8)
	m.Push(3, 7)
	m.Push(4, 6)
	x.ExpectNotOk(m.At(4))
	x.ExpectOk(m.At(2)).ToBe(7)
	x.ExpectOk(m.At(3)).ToBe(6)
}

func TestOMRange(t *testing.T) {
	x := expected.New(t)
	m := NewOrderedMap[int, int]()
	m.Push(1, 9)
	m.Push(2, 8)
	m.Push(3, 7)
	keys := []int{}
	values := []int{}
	m.Range(func(k int, v int) {
		keys = append(keys, k)
		values = append(values, v)
	})
	x.Expect(keys).ToBe([]int{1, 2, 3})
	x.Expect(values).ToBe([]int{9, 8, 7})
}
