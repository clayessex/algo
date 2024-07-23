package vessels

import (
	"testing"
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
	s := NewStack[int](4)
	s.Push(9)
	s.Push(8)
	s.Push(7)
	expect(t, s.Len(), 3)
	expect(t, s.Pop(), 7)
	expect(t, s.Pop(), 8)
	expect(t, s.Pop(), 9)
	expect(t, s.Len(), 0)

	// detect panic
	defer func() { _ = recover() }()
	s.Pop() // Should panic
	t.Fatal("Stack Pop should panic when empty")
}

func TestStackAt(t *testing.T) {
	s := NewStack[int](5)
	s.Push(9)
	s.Push(8)
	s.Push(7)
	s.Push(6)
	expect(t, s.At(0), 9)
	expect(t, s.Pop(), 6)
	expect(t, s.At(0), 9)
	expect(t, s.At(1), 8)
	expect(t, s.At(2), 7)

	defer func() { _ = recover() }()
	s.At(3)
	t.Fatal("Queue At() with an invalid index should have panicked")
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
	s := NewStack[int](4)
	s.Push(9)
	s.Push(8)
	s.Push(7)
	expect(t, s.Len(), 3)
	c := s.Clone()
	expect(t, c.Len(), 3)
	expect(t, c.Pop(), 7)
	expect(t, c.Pop(), 8)
	expect(t, c.Pop(), 9)
	expect(t, c.Len(), 0)
	expect(t, s.Len(), 3)
}
