package vessels

import (
	"testing"

	"github.com/clayessex/algo/expected"
)

func TestNewStack(t *testing.T) {
	s := NewStack[int](2)
	expect(t, s.Len(), 0)
	expect(t, s.Cap(), 2)
}

func TestStackLen(t *testing.T) {
	s := NewStack[int]()
	expect(t, s.Len(), 0)
	s.Push(9)
	expect(t, s.Len(), 1)
	s.Push(9)
	expect(t, s.Len(), 2)
	s.Pop()
	s.Pop()
	expect(t, s.Len(), 0)
}

func TestStackCap(t *testing.T) {
	s := NewStack[int](2)
	expect(t, s.Cap(), 2)
	s.Push(9)
	s.Push(8)
	expect(t, s.Cap(), 2) // warn: internal
	s.Push(7)
	expect(t, s.Cap(), 4) // warn: internal
}

func TestStackPushPop(t *testing.T) {
	x := expected.New(t)
	s := NewStack[int](4)
	s.Push(9)
	s.Push(8)
	s.Push(7)
	x.Expect(s.Len()).ToBe(3)
	x.ExpectOk(s.Pop()).ToBe(7)
	x.ExpectOk(s.Pop()).ToBe(8)
	x.ExpectOk(s.Pop()).ToBe(9)
	x.Expect(s.Len()).ToBe(0)
	x.ExpectNotOk(s.Pop())
}

func TestStackAt(t *testing.T) {
	x := expected.New(t)
	s := NewStack[int](5)
	s.Push(9)
	s.Push(8)
	s.Push(7)
	s.Push(6)
	x.ExpectOk(s.At(0)).ToBe(9)
	x.ExpectOk(s.Pop()).ToBe(6)
	x.ExpectOk(s.At(0)).ToBe(9)
	x.ExpectOk(s.At(1)).ToBe(8)
	x.ExpectOk(s.At(2)).ToBe(7)
	x.ExpectNotOk(s.At(3))
}

func TestStackClear(t *testing.T) {
	s := NewStack[int](4)
	s.Push(9)
	s.Push(8)
	s.Push(7)
	expect(t, s.Len(), 3)
	s.Clear()
	expect(t, s.Len(), 0)
}

func TestStackClone(t *testing.T) {
	x := expected.New(t)
	s := NewStack[int](4)
	s.Push(9)
	s.Push(8)
	s.Push(7)
	x.Expect(s.Len()).ToBe(3)
	c := s.Clone()
	x.Expect(c.Len()).ToBe(3)
	x.ExpectOk(c.Pop()).ToBe(7)
	x.ExpectOk(c.Pop()).ToBe(8)
	x.ExpectOk(c.Pop()).ToBe(9)
	x.Expect(c.Len()).ToBe(0)
	x.Expect(s.Len()).ToBe(3)
}
