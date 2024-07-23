package vessels

import (
	"testing"
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
	d := NewDeque[int](10)
	d.PushBack(9)
	d.PushBack(8)
	d.PushBack(7)
	d.PushBack(6)
	expect(t, d.Len(), 4)

	b := make([]int, 4)
	d.copy(b)
	expect(t, b, []int{9, 8, 7, 6})

	// Test expected panic
	defer func() { _ = recover() }()
	b = make([]int, 2)
	// Copy to an undersized buffer panics
	d.copy(b)
	t.Fatal("Deque copy() to a smaller buffer should have panicked")
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
	d := NewDeque[int](4)
	d.PushBack(9)
	expect(t, d.Len(), 1)
	d.PushBack(8)
	d.PushBack(7)
	expect(t, d.Len(), 3)
	expect(t, d.PopBack(), 7)
	expect(t, d.PopBack(), 8)
	expect(t, d.PopBack(), 9)
	expect(t, d.Len(), 0)
	d.PushBack(6)
	d.PushBack(5)
	d.PushBack(4)
	d.PushBack(3) // resize
	d.PushBack(2)
	expect(t, d.Len(), 5)
}

func TestPushFront(t *testing.T) {
	d := NewDeque[int](4)
	d.PushFront(9)
	expect(t, d.Len(), 1)
	expect(t, d.PopFront(), 9)

	d.PushFront(8)
	d.PushFront(7)
	d.PushFront(6)
	expect(t, d.Len(), 3)
	expect(t, d.PopFront(), 6)
	expect(t, d.PopBack(), 8)
	expect(t, d.PopFront(), 7)

	d.PushFront(5)
	d.PushFront(4)
	d.PushFront(3)
	d.PushFront(2) // resize
	d.PushFront(1)
	expect(t, d.Len(), 5)
}

func TestPopBack(t *testing.T) {
	d := NewDeque[int](4)
	d.PushBack(9)
	expect(t, d.Len(), 1)
	expect(t, d.PopBack(), 9)
	expect(t, d.Len(), 0)
	d.PushBack(7)
	d.PushBack(6)
	d.PushFront(8)
	expect(t, d.Len(), 3)
	expect(t, d.PopBack(), 6)
	expect(t, d.PopBack(), 7)
	expect(t, d.PopBack(), 8)
	expect(t, d.Len(), 0)

	// Test panic
	defer func() { _ = recover() }()
	expect(t, d.Len(), 0)
	d.PopBack()
	t.Fatal("PopBack on an empty deque should have panicked")
}

func TestPopFront(t *testing.T) {
	d := NewDeque[int](4)
	d.PushBack(9)
	expect(t, d.Len(), 1)
	expect(t, d.PopFront(), 9)
	expect(t, d.Len(), 0)
	d.PushBack(7)
	d.PushBack(6)
	d.PushFront(8)
	expect(t, d.Len(), 3)
	expect(t, d.PopFront(), 8)
	expect(t, d.PopFront(), 7)
	expect(t, d.PopFront(), 6)
	expect(t, d.Len(), 0)

	// Test panic
	defer func() { _ = recover() }()
	expect(t, d.Len(), 0)
	d.PopFront()
	t.Fatal("PopFront on an empty deque should have panicked")
}

func TestDequeFront(t *testing.T) {
	d := NewDeque[int](4)
	d.PushBack(9)
	d.PushBack(8)
	d.PushBack(7)
	expect(t, d.Len(), 3)
	expect(t, d.Front(), 9)
	expect(t, d.Len(), 3)

	expect(t, d.PopBack(), 7)
	expect(t, d.PopBack(), 8)
	expect(t, d.PopBack(), 9)

	defer func() { _ = recover() }()
	d.Front()
	t.Fatal("Front() on empty Deque should panic")
}

func TestDequeBack(t *testing.T) {
	d := NewDeque[int](4)
	d.PushBack(9)
	d.PushBack(8)
	d.PushBack(7)
	expect(t, d.Len(), 3)
	expect(t, d.Back(), 7)
	expect(t, d.Len(), 3)

	expect(t, d.PopBack(), 7)
	expect(t, d.PopBack(), 8)
	expect(t, d.PopBack(), 9)

	defer func() { _ = recover() }()
	d.Back()
	t.Fatal("Back() on empty Deque should panic")
}

func TestAt(t *testing.T) {
	d := NewDeque[int](4)
	d.PushBack(9)
	d.PushBack(8)
	d.PushBack(7)
	d.PopFront()
	expect(t, d.At(0), 8)
	expect(t, d.At(1), 7)
	defer func() { _ = recover() }()
	d.At(2)
	t.Fatal("deque At() with an invalid index should have panicked")
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
	d := NewDeque[int](4)
	d.PushBack(9)
	d.PushBack(8)
	d.PushBack(7)
	expect(t, d.Len(), 3)
	c := d.Clone()
	expect(t, c.Len(), 3)
	expect(t, c.PopBack(), 7)
	expect(t, c.PopBack(), 8)
	expect(t, c.PopBack(), 9)
	expect(t, c.Len(), 0)
	expect(t, d.Len(), 3)

	// Test buffer equality and Deque.copy()
	d = NewDeque[int](8)
	d.PushBack(9)
	d.PushBack(8)
	d.PushBack(7)
	d.PushBack(6)
	d.PushBack(5)
	c = d.Clone()

	expect(t, d.Len(), c.Len())
	for i := 0; i < d.Len(); i++ {
		expect(t, d.PopBack(), c.PopBack())
	}
}

func TestAutoResize(t *testing.T) {
	d := NewDeque[int](2)
	d.PushBack(9)
	d.PushBack(8)
	expect(t, d.Len(), 2)
	expect(t, d.Cap(), 2) // Internal
	d.PushBack(7)
	d.PushBack(6)
	d.PushBack(5)
	expect(t, d.Len(), 5)
	expect(t, d.Cap(), 8) // Internal

	expect(t, d.PopBack(), 5)
	expect(t, d.PopBack(), 6)
	expect(t, d.PopBack(), 7)
	expect(t, d.PopBack(), 8)
	expect(t, d.PopBack(), 9)

	expect(t, d.Len(), 0)
	expect(t, d.Cap(), 8) // Internal
}

func TestGenerics(t *testing.T) {
	a := NewDeque[string](2)
	a.PushBack("test-a1")
	a.PushBack("test-a2")
	expect(t, a.PopBack(), "test-a2")
	expect(t, a.PopBack(), "test-a1")
	b := NewDeque[float64](2)
	b.PushBack(3.1415927)
	b.PushBack(2.71828)
	expect(t, b.PopBack(), 2.71828)   // warn: float-cmp
	expect(t, b.PopBack(), 3.1415927) // warn: float-cmp
	c := NewDeque[byte](2)
	c.PushBack(byte(13))
	c.PushBack(byte(10))
	expect(t, c.PopBack(), byte(10))
	expect(t, c.PopBack(), byte(13))
}

func TestShrink(t *testing.T) {
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

	expect(t, d.PopBack(), 7)
	expect(t, d.PopBack(), 8)
	expect(t, d.PopBack(), 9)
}
