package vessels

import (
	"testing"

	"github.com/clayessex/algo/expected"
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
	expect(t, q.Cap(), 2) // warn: internal
	q.Push(7)
	expect(t, q.Cap(), 4) // warn: internal
}

func TestQueuePushPop(t *testing.T) {
	x := expected.New(t)
	q := NewQueue[int](4)
	q.Push(9)
	q.Push(8)
	q.Push(7)
	x.Expect(q.Len()).ToBe(3)
	x.ExpectOk(q.Pop()).ToBe(9)
	x.ExpectOk(q.Pop()).ToBe(8)
	x.ExpectOk(q.Pop()).ToBe(7)
	x.Expect(q.Len()).ToBe(0)
	x.ExpectNotOk(q.Pop())
}

func TestQueueAt(t *testing.T) {
	x := expected.New(t)
	q := NewQueue[int](5)
	q.Push(9)
	q.Push(8)
	q.Push(7)
	q.Push(6)
	x.ExpectOk(q.At(0)).ToBe(9)
	x.ExpectOk(q.Pop()).ToBe(9)
	x.ExpectOk(q.At(0)).ToBe(8)
	x.ExpectOk(q.At(1)).ToBe(7)
	x.ExpectOk(q.At(2)).ToBe(6)
	x.ExpectNotOk(q.At(3))
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
	x := expected.New(t)
	d := NewQueue[int](4)
	d.Push(9)
	d.Push(8)
	d.Push(7)
	expect(t, d.Len(), 3)
	c := d.Clone()
	x.Expect(c.Len()).ToBe(3)
	x.ExpectOk(c.Pop()).ToBe(9)
	x.ExpectOk(c.Pop()).ToBe(8)
	x.ExpectOk(c.Pop()).ToBe(7)
	x.Expect(c.Len()).ToBe(0)
	x.Expect(d.Len()).ToBe(3)
}
