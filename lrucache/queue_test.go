package main

import (
	"testing"
)

func TestMakeQueueSized(t *testing.T) {
	q := MakeQueueSized[int](2)
	expect(t, q.Size(), 0)
}

func TestMakeQueue(t *testing.T) {
	if q := MakeQueue[int](); q.Cap() <= 0 {
		t.Fatalf("MakeQueue failed with size: %v\n", q.Cap())
	}
}

func TestQueueSize(t *testing.T) {
	q := MakeQueue[int]()
	expect(t, q.Size(), 0)
	q.Push(9)
	expect(t, q.Size(), 1)
	q.Push(9)
	expect(t, q.Size(), 2)
	q.Pop()
	q.Pop()
	expect(t, q.Size(), 0)
}

func TestQueueCap(t *testing.T) {
	q := MakeQueueSized[int](2)
	expect(t, q.Cap(), 2)
	q.Push(9)
	q.Push(8)
	expect(t, q.Cap(), 4) // warn: internal
}

func TestPush(t *testing.T) {
	q := (*Queue[int])(MakeDequeSized[int](2))
	q.Push(9)
	expect(t, q.Size(), 1)
	q.Push(8)
	q.Pop()
	expect(t, q.Size(), 1)
}

func TestPop(t *testing.T) {
	q := (*Queue[int])(MakeDequeSized[int](2))
	q.Push(9)
	q.Push(8)
	q.Push(7)
	expect(t, q.Pop(), 9)
	expect(t, q.Pop(), 8)
	expect(t, q.Pop(), 7)

	// detect panic
	defer func() { _ = recover() }()
	q.Pop()
	t.Fatal("Queue Pop should have panicked")
}

func TestQueueClear(t *testing.T) {
	d := MakeQueueSized[int](4)
	d.Push(9)
	d.Push(8)
	d.Push(7)
	expect(t, d.Size(), 3)
	d.Clear()
	expect(t, d.Size(), 0)
}

func TestQueueClone(t *testing.T) {
	d := MakeQueueSized[int](4)
	d.Push(9)
	d.Push(8)
	d.Push(7)
	expect(t, d.Size(), 3)
	c := d.Clone()
	expect(t, c.Size(), 3)
	expect(t, c.Pop(), 9)
	expect(t, c.Pop(), 8)
	expect(t, c.Pop(), 7)
	expect(t, c.Size(), 0)
	expect(t, d.Size(), 3)
}
