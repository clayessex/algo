package vessels

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[int](2)
	expect(t, q.Cap(), 2)
	expect(t, q.Len(), 0)
}

func TestQueueLen(t *testing.T) {
	q := NewQueue[int]()
	expect(t, q.Len(), 0)
	q.Push(9)
	expect(t, q.Len(), 1)
	q.Push(9)
	expect(t, q.Len(), 2)
	q.Pop()
	q.Pop()
	expect(t, q.Len(), 0)
}

func TestQueueCap(t *testing.T) {
	q := NewQueue[int](2)
	expect(t, q.Cap(), 2)
	q.Push(9)
	q.Push(8)
	expect(t, q.Cap(), 4) // warn: internal
}

func TestQueuePushPop(t *testing.T) {
	q := NewQueue[int](4)
	q.Push(9)
	q.Push(8)
	q.Push(7)
	expect(t, q.Len(), 3)
	expect(t, q.Pop(), 9)
	expect(t, q.Pop(), 8)
	expect(t, q.Pop(), 7)
	expect(t, q.Len(), 0)

	// detect panic
	defer func() { _ = recover() }()
	q.Pop() // Should panic
	t.Fatal("Queue Pop should panic when empty")
}

func TestQueueAt(t *testing.T) {
	q := NewQueue[int](5)
	q.Push(9)
	q.Push(8)
	q.Push(7)
	q.Push(6)
	expect(t, q.At(0), 9)
	expect(t, q.Pop(), 9)
	expect(t, q.At(0), 8)
	expect(t, q.At(1), 7)
	expect(t, q.At(2), 6)

	defer func() { _ = recover() }()
	q.At(3)
	t.Fatal("Queue At() with an invalid index should have panicked")
}

func TestQueueClear(t *testing.T) {
	d := NewQueue[int](4)
	d.Push(9)
	d.Push(8)
	d.Push(7)
	expect(t, d.Len(), 3)
	d.Clear()
	expect(t, d.Len(), 0)
}

func TestQueueClone(t *testing.T) {
	d := NewQueue[int](4)
	d.Push(9)
	d.Push(8)
	d.Push(7)
	expect(t, d.Len(), 3)
	c := d.Clone()
	expect(t, c.Len(), 3)
	expect(t, c.Pop(), 9)
	expect(t, c.Pop(), 8)
	expect(t, c.Pop(), 7)
	expect(t, c.Len(), 0)
	expect(t, d.Len(), 3)
}
