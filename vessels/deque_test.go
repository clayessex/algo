package vessels

import (
	"testing"

	"github.com/clayessex/algo/expected"
)

func dequeBufferContains[T comparable](d *Deque[T], want []T) bool {
	index := d.tail
	for _, v := range want {
		if v != d.buf[index] {
			return false
		}
		index = (index + 1) % len(d.buf)
	}
	return true
}

func makeTestDeque[T any](values []T, head int, tail int) *Deque[T] {
	return &Deque[T]{values, head, tail}
}

func Test_internal_NewDeque(t *testing.T) {
	d := NewDeque[int](4)
	if len(d.buf) != 5 {
		t.Fatalf("NewDeque failed: expected cap: %v, got %v", 5, len(d.buf))
	}
	if len(d.buf) != 5 {
		t.Fatalf("NewDeque failed: expected len: %v, got %v", 5, len(d.buf))
	}
}

func Test_internal_DequeSize(t *testing.T) {
	d := &Deque[int]{[]int{0, 0}, 0, 0}
	expect(t, d.Len(), 0)
	d = &Deque[int]{[]int{9, 0}, 1, 0}
	expect(t, d.Len(), 1)
	d = &Deque[int]{[]int{9, 8, 0, 0}, 2, 0}
	expect(t, d.Len(), 2)
	d = &Deque[int]{[]int{0, 9, 8, 0}, 3, 1}
	expect(t, d.Len(), 2)
	d = &Deque[int]{[]int{8, 7, 0, 9}, 2, 3}
	expect(t, d.Len(), 3)
}

func Test_internal_copy(t *testing.T) {
	x := expected.New(t)
	d := NewDeque[int](10)
	d.PushBack(9)
	d.PushBack(8)
	d.PushBack(7)
	d.PushBack(6)
	x.Expect(d.Len()).ToBe(4)

	b := make([]int, 4)
	d.copy(b)
	x.Expect(b).ToBe([]int{9, 8, 7, 6})

	// Copy to an undersized buffer
	b = make([]int, 2)
	x.Expect(d.Len()).ToBe(4)
	x.Expect(d.copy(b)).ToBe(2)
}

func Test_internal_PushBack(t *testing.T) {
	d := &Deque[int]{[]int{0, 0}, 0, 0}
	d.PushBack(9)
	expect(t, d, &Deque[int]{[]int{9, 0}, 1, 0})
	d.PushBack(8)
	expect(t, d, &Deque[int]{[]int{9, 8, 0}, 2, 0})
	d.PushBack(7)
	expect(t, d, &Deque[int]{[]int{9, 8, 7, 0, 0}, 3, 0})

	d = &Deque[int]{[]int{0, 8, 7, 0}, 3, 1}
	d.PushBack(6)
	expect(t, d, &Deque[int]{[]int{0, 8, 7, 6}, 0, 1})
	d.PushBack(5)
	expect(t, d, &Deque[int]{[]int{8, 7, 6, 5, 0, 0, 0}, 4, 0})

	d = &Deque[int]{[]int{7, 0, 0, 8}, 1, 3}
	d.PushBack(6)
	expect(t, d, &Deque[int]{[]int{7, 6, 0, 8}, 2, 3})
}

func TestNewDeque(t *testing.T) {
	d := NewDeque[int](4)
	expect(t, d.Cap(), 4)
	expect(t, d.Len(), 0)
	d = NewDeque[int](32)
	expect(t, d.Cap(), 32)
	expect(t, d.Len(), 0)
	d = NewDeque[int](1000)
	expect(t, d.Cap(), 1000)
	expect(t, d.Len(), 0)
}

func TestMakeDeque(t *testing.T) {
	d := NewDeque[int]()
	if d.Cap() <= 0 { // Internal default larger than zero
		t.Fatalf("expected: >0, got: %v\n", d.Cap())
	}
}

func TestSize(t *testing.T) {
	d := NewDeque[int](4)
	expect(t, d.Len(), 0)
	d.PushBack(9)
	expect(t, d.Len(), 1)
	d.PopBack()
	expect(t, d.Len(), 0)
	d.PushBack(0)
	d.PushBack(1)
	d.PopFront()
	d.PopFront()

	d.PushBack(2)
	expect(t, d.Len(), 1)
	d.PushBack(3)
	expect(t, d.Len(), 2)
	d.PushBack(4)
	expect(t, d.Len(), 3)
	d.PopFront()
	expect(t, d.Len(), 2)
	d.PopFront()
	expect(t, d.Len(), 1)
	d.PopFront()
	expect(t, d.Len(), 0)
	d.PushBack(3)
	d.PushBack(4)
	d.PushBack(5)
	expect(t, d.Len(), 3)
	d.PopFront()
	d.PopFront()
	d.PopFront()
	expect(t, d.Len(), 0)
}

func TestCap(t *testing.T) {
	d := NewDeque[int](2)
	expect(t, d.Cap(), 2)
	d.PushBack(9)
	d.PushBack(8)
	expect(t, d.Cap(), 2)
	d.PushBack(8)
	expect(t, d.Cap(), 4)
}

func TestPushBack(t *testing.T) {
	x := expected.New(t)
	d := NewDeque[int](4)
	d.PushBack(9)
	x.Expect(d.Len()).ToBe(1)
	d.PushBack(8)
	d.PushBack(7)
	x.Expect(d.Len()).ToBe(3)
	x.ExpectOk(d.PopBack()).ToBe(7)
	x.ExpectOk(d.PopBack()).ToBe(8)
	x.ExpectOk(d.PopBack()).ToBe(9)
	x.Expect(d.Len()).ToBe(0)
	d.PushBack(6)
	d.PushBack(5)
	d.PushBack(4)
	d.PushBack(3) // resize
	d.PushBack(2)
	x.Expect(d.Len()).ToBe(5)
}

func TestPushFront(t *testing.T) {
	x := expected.New(t)
	d := NewDeque[int](4)
	d.PushFront(9)
	x.Expect(d.Len()).ToBe(1)
	x.ExpectOk(d.PopFront()).ToBe(9)

	d.PushFront(8)
	d.PushFront(7)
	d.PushFront(6)
	x.Expect(d.Len()).ToBe(3)
	x.ExpectOk(d.PopFront()).ToBe(6)
	x.ExpectOk(d.PopBack()).ToBe(8)
	x.ExpectOk(d.PopFront()).ToBe(7)

	d.PushFront(5)
	d.PushFront(4)
	d.PushFront(3)
	d.PushFront(2) // resize
	d.PushFront(1)
	x.Expect(d.Len()).ToBe(5)
}

func TestPopBack(t *testing.T) {
	x := expected.New(t)
	d := NewDeque[int](4)
	d.PushBack(9)
	x.Expect(d.Len()).ToBe(1)
	x.ExpectOk(d.PopBack()).ToBe(9)
	x.Expect(d.Len()).ToBe(0)
	d.PushBack(7)
	d.PushBack(6)
	d.PushFront(8)
	x.Expect(d.Len()).ToBe(3)
	x.ExpectOk(d.PopBack()).ToBe(6)
	x.ExpectOk(d.PopBack()).ToBe(7)
	x.ExpectOk(d.PopBack()).ToBe(8)
	x.ExpectNotOk(d.PopBack()) // on empty
}

func TestPopFront(t *testing.T) {
	x := expected.New(t)
	d := NewDeque[int](4)
	d.PushBack(9)
	x.Expect(d.Len()).ToBe(1)
	x.ExpectOk(d.PopFront()).ToBe(9)
	x.Expect(d.Len()).ToBe(0)
	d.PushBack(7)
	d.PushBack(6)
	d.PushFront(8)
	x.Expect(d.Len()).ToBe(3)
	x.ExpectOk(d.PopFront()).ToBe(8)
	x.ExpectOk(d.PopFront()).ToBe(7)
	x.ExpectOk(d.PopFront()).ToBe(6)
	x.Expect(d.Len()).ToBe(0)
	x.ExpectNotOk(d.PopFront()) // on empty
}

func TestDequeFront(t *testing.T) {
	x := expected.New(t)
	d := NewDeque[int](4)
	d.PushBack(9)
	d.PushBack(8)
	d.PushBack(7)
	x.Expect(d.Len()).ToBe(3)
	x.ExpectOk(d.Front()).ToBe(9)
	x.Expect(d.Len()).ToBe(3)
	x.ExpectOk(d.PopBack()).ToBe(7)
	x.ExpectOk(d.PopBack()).ToBe(8)
	x.ExpectOk(d.PopBack()).ToBe(9)
	x.ExpectNotOk(d.Front()) // on empty
}

func TestDequeBack(t *testing.T) {
	x := expected.New(t)
	d := NewDeque[int](4)
	d.PushBack(9)
	d.PushBack(8)
	d.PushBack(7)
	x.Expect(d.Len()).ToBe(3)
	x.ExpectOk(d.Back()).ToBe(7)
	x.Expect(d.Len()).ToBe(3)
	x.ExpectOk(d.PopBack()).ToBe(7)
	x.ExpectOk(d.PopBack()).ToBe(8)
	x.ExpectOk(d.PopBack()).ToBe(9)
	x.ExpectNotOk(d.Back()) // on empty
}

func TestAt(t *testing.T) {
	x := expected.New(t)
	d := NewDeque[int](4)
	d.PushBack(9)
	d.PushBack(8)
	d.PushBack(7)
	d.PopFront()
	x.ExpectOk(d.At(0)).ToBe(8)
	x.ExpectOk(d.At(1)).ToBe(7)
	x.ExpectNotOk(d.At(2))
}

func TestClear(t *testing.T) {
	d := NewDeque[int](4)
	d.PushBack(9)
	d.PushBack(8)
	d.PushBack(7)
	expect(t, d.Len(), 3)
	d.Clear()
	expect(t, d.Len(), 0)
}

func TestClone(t *testing.T) {
	x := expected.New(t)
	d := NewDeque[int](4)
	d.PushBack(9)
	d.PushBack(8)
	d.PushBack(7)
	x.Expect(d.Len()).ToBe(3)
	c := d.Clone()
	x.Expect(c.Len()).ToBe(3)
	x.ExpectOk(c.PopBack()).ToBe(7)
	x.ExpectOk(c.PopBack()).ToBe(8)
	x.ExpectOk(c.PopBack()).ToBe(9)
	x.Expect(c.Len()).ToBe(0)
	x.Expect(d.Len()).ToBe(3)

	// Test buffer equality and Deque.copy()
	d = NewDeque[int](8)
	d.PushBack(9)
	d.PushBack(8)
	d.PushBack(7)
	d.PushBack(6)
	d.PushBack(5)
	c = d.Clone()

	x.Expect(d.Len()).ToBe(c.Len())
	for i := 0; i < d.Len(); i++ {
		value := x.ExpectOk(c.PopBack()).Value()
		x.ExpectOk(d.PopBack()).ToBe(value)
	}
}

func TestAutoResize(t *testing.T) {
	x := expected.New(t)
	d := NewDeque[int](2)
	d.PushBack(9)
	d.PushBack(8)
	x.Expect(d.Len()).ToBe(2)
	x.Expect(d.Cap()).ToBe(2) // Internal
	d.PushBack(7)
	d.PushBack(6)
	d.PushBack(5)
	x.Expect(d.Len()).ToBe(5)
	x.Expect(d.Cap()).ToBe(8) // Internal

	x.ExpectOk(d.PopBack()).ToBe(5)
	x.ExpectOk(d.PopBack()).ToBe(6)
	x.ExpectOk(d.PopBack()).ToBe(7)
	x.ExpectOk(d.PopBack()).ToBe(8)
	x.ExpectOk(d.PopBack()).ToBe(9)

	x.Expect(d.Len()).ToBe(0)
	x.Expect(d.Cap()).ToBe(8) // Internal
}

func TestGenerics(t *testing.T) {
	x := expected.New(t)
	a := NewDeque[string](2)
	a.PushBack("test-a1")
	a.PushBack("test-a2")
	x.ExpectOk(a.PopBack()).ToBe("test-a2")
	x.ExpectOk(a.PopBack()).ToBe("test-a1")
	b := NewDeque[float64](2)
	b.PushBack(3.1415927)
	b.PushBack(2.71828)
	x.ExpectOk(b.PopBack()).ToBe(2.71828)   // warn: float-cmp
	x.ExpectOk(b.PopBack()).ToBe(3.1415927) // warn: float-cmp
	c := NewDeque[byte](2)
	c.PushBack(byte(13))
	c.PushBack(byte(10))
	x.ExpectOk(c.PopBack()).ToBe(byte(10))
	x.ExpectOk(c.PopBack()).ToBe(byte(13))
}

func TestShrink(t *testing.T) {
	x := expected.New(t)
	d := NewDeque[int](128)
	d.PushBack(9)
	d.PushBack(8)
	d.PushBack(7)

	d.Shrink()
	expect(t, d.Cap(), 64)
	d.Shrink()
	expect(t, d.Cap(), 32)
	d.Shrink()
	expect(t, d.Cap(), 32) // Internal (min default initial size)

	x.ExpectOk(d.PopBack()).ToBe(7)
	x.ExpectOk(d.PopBack()).ToBe(8)
	x.ExpectOk(d.PopBack()).ToBe(9)
}
