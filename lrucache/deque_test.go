package main

import (
	"reflect"
	"testing"
)

func dequeBufferContains[T comparable](d *Deque[T], want []T) bool {
	index := d.tail
	for _, v := range want {
		if v != d.buf[index] {
			return false
		}
		index = (index + 1) % d.cap
	}
	return true
}

func makeTestDeque[T any](values []T, head int, tail int) *Deque[T] {
	return &Deque[T]{values, len(values), head, tail}
}

func expect(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("%s failed:\n    expected: %v\n      actual: %v\n", t.Name(), expected, actual)
	}
}

func Test_internal_MakeDequeSized(t *testing.T) {
	d := MakeDequeSized[int](4)
	if d.cap != 4 {
		t.Fatalf("MakeDequeSized failed: expected cap: %v, got %v", 4, d.cap)
	}
	if len(d.buf) != 4 {
		t.Fatalf("MakeDequeSized failed: expected len: %v, got %v", 4, len(d.buf))
	}
}

func Test_internal_DequeSize(t *testing.T) {
	d := &Deque[int]{[]int{0, 0}, 2, 0, 0}
	expect(t, d.Size(), 0)
	d = &Deque[int]{[]int{9, 0}, 2, 1, 0}
	expect(t, d.Size(), 1)
	d = &Deque[int]{[]int{9, 8, 0, 0}, 4, 2, 0}
	expect(t, d.Size(), 2)
	d = &Deque[int]{[]int{0, 9, 8, 0}, 4, 3, 1}
	expect(t, d.Size(), 2)
	d = &Deque[int]{[]int{8, 7, 0, 9}, 4, 2, 3}
	expect(t, d.Size(), 3)
}

func Test_internal_PushBack(t *testing.T) {
	d := &Deque[int]{[]int{0, 0}, 2, 0, 0}
	d.PushBack(9)
	expect(t, d, &Deque[int]{[]int{9, 0}, 2, 1, 0})
	d.PushBack(8)
	expect(t, d, &Deque[int]{[]int{9, 8, 0, 0}, 4, 2, 0})
	d.PushBack(7)
	expect(t, d, &Deque[int]{[]int{9, 8, 7, 0}, 4, 3, 0})

	d = &Deque[int]{[]int{0, 8, 7, 0}, 4, 3, 1}
	d.PushBack(6)
	expect(t, d, &Deque[int]{[]int{0, 8, 7, 6}, 4, 0, 1})
	d.PushBack(5)
	expect(t, d, &Deque[int]{[]int{8, 7, 6, 5, 0, 0, 0, 0}, 8, 4, 0})

	d = &Deque[int]{[]int{7, 0, 0, 8}, 4, 1, 3}
	d.PushBack(6)
	expect(t, d, &Deque[int]{[]int{7, 6, 0, 8}, 4, 2, 3})
}

func TestMakeDequeSized(t *testing.T) {
	d := MakeDequeSized[int](4)
	expect(t, d.Cap(), 4)
	expect(t, d.Size(), 0)
	d = MakeDequeSized[int](32)
	expect(t, d.Cap(), 32)
	expect(t, d.Size(), 0)
	d = MakeDequeSized[int](1000)
	expect(t, d.Cap(), 1000)
	expect(t, d.Size(), 0)
}

func TestMakeDeque(t *testing.T) {
	d := MakeDeque[int]()
	if d.Cap() <= 0 { // Internal default larger than zero
		t.Fatalf("expected: >0, got: %v\n", d.Cap())
	}
}

func TestSize(t *testing.T) {
	d := MakeDequeSized[int](4)
	expect(t, d.Size(), 0)
	d.PushBack(9)
	expect(t, d.Size(), 1)
	d.PopBack()
	expect(t, d.Size(), 0)
	d.PushBack(0)
	d.PushBack(1)
	d.PopFront()
	d.PopFront()

	d.PushBack(2)
	expect(t, d.Size(), 1)
	d.PushBack(3)
	expect(t, d.Size(), 2)
	d.PushBack(4)
	expect(t, d.Size(), 3)
	d.PopFront()
	expect(t, d.Size(), 2)
	d.PopFront()
	expect(t, d.Size(), 1)
	d.PopFront()
	expect(t, d.Size(), 0)
	d.PushBack(3)
	d.PushBack(4)
	d.PushBack(5)
	expect(t, d.Size(), 3)
	d.PopFront()
	d.PopFront()
	d.PopFront()
	expect(t, d.Size(), 0)
}

func TestCap(t *testing.T) {
	d := MakeDequeSized[int](2)
	expect(t, d.Cap(), 2)
	d.PushBack(9)
	d.PushBack(8)
	expect(t, d.Cap(), 4)
}

func TestPushBack(t *testing.T) {
	d := MakeDequeSized[int](4)
	d.PushBack(9)
	expect(t, d.Size(), 1)
	d.PushBack(8)
	d.PushBack(7)
	expect(t, d.Size(), 3)
	expect(t, d.PopBack(), 7)
	expect(t, d.PopBack(), 8)
	expect(t, d.PopBack(), 9)
	expect(t, d.Size(), 0)
	d.PushBack(6)
	d.PushBack(5)
	d.PushBack(4)
	d.PushBack(3) // resize
	d.PushBack(2)
	expect(t, d.Size(), 5)
}

func TestPushFront(t *testing.T) {
	d := MakeDequeSized[int](4)
	d.PushFront(9)
	expect(t, d.Size(), 1)
	expect(t, d.PopFront(), 9)

	d.PushFront(8)
	d.PushFront(7)
	d.PushFront(6)
	expect(t, d.Size(), 3)
	expect(t, d.PopFront(), 6)
	expect(t, d.PopBack(), 8)
	expect(t, d.PopFront(), 7)

	d.PushFront(5)
	d.PushFront(4)
	d.PushFront(3)
	d.PushFront(2) // resize
	d.PushFront(1)
	expect(t, d.Size(), 5)
}

func TestPopBack(t *testing.T) {
	d := MakeDequeSized[int](4)
	d.PushBack(9)
	expect(t, d.Size(), 1)
	expect(t, d.PopBack(), 9)
	expect(t, d.Size(), 0)
	d.PushBack(7)
	d.PushBack(6)
	d.PushFront(8)
	expect(t, d.Size(), 3)
	expect(t, d.PopBack(), 6)
	expect(t, d.PopBack(), 7)
	expect(t, d.PopBack(), 8)
	expect(t, d.Size(), 0)

	// Test panic
	defer func() { _ = recover() }()
	expect(t, d.Size(), 0)
	d.PopBack()
	t.Fatal("PopBack on an empty deque should have panicked")
}

func TestPopFront(t *testing.T) {
	d := MakeDequeSized[int](4)
	d.PushBack(9)
	expect(t, d.Size(), 1)
	expect(t, d.PopFront(), 9)
	expect(t, d.Size(), 0)
	d.PushBack(7)
	d.PushBack(6)
	d.PushFront(8)
	expect(t, d.Size(), 3)
	expect(t, d.PopFront(), 8)
	expect(t, d.PopFront(), 7)
	expect(t, d.PopFront(), 6)
	expect(t, d.Size(), 0)

	// Test panic
	defer func() { _ = recover() }()
	expect(t, d.Size(), 0)
	d.PopFront()
	t.Fatal("PopFront on an empty deque should have panicked")
}

func TestClear(t *testing.T) {
	d := MakeDequeSized[int](4)
	d.PushBack(9)
	d.PushBack(8)
	d.PushBack(7)
	expect(t, d.Size(), 3)
	d.Clear()
	expect(t, d.Size(), 0)
}

func TestClone(t *testing.T) {
	d := MakeDequeSized[int](4)
	d.PushBack(9)
	d.PushBack(8)
	d.PushBack(7)
	expect(t, d.Size(), 3)
	c := d.Clone()
	expect(t, c.Size(), 3)
	expect(t, c.PopBack(), 7)
	expect(t, c.PopBack(), 8)
	expect(t, c.PopBack(), 9)
	expect(t, c.Size(), 0)
	expect(t, d.Size(), 3)
}

func TestAutoResize(t *testing.T) {
	d := MakeDequeSized[int](2)
	d.PushBack(9)
	d.PushBack(8)
	expect(t, d.Size(), 2)
	expect(t, d.Cap(), 4) // Internal
	d.PushBack(7)
	d.PushBack(6)
	d.PushBack(5)
	expect(t, d.Size(), 5)
	expect(t, d.Cap(), 8) // Internal

	expect(t, d.PopBack(), 5)
	expect(t, d.PopBack(), 6)
	expect(t, d.PopBack(), 7)
	expect(t, d.PopBack(), 8)
	expect(t, d.PopBack(), 9)

	expect(t, d.Size(), 0)
	expect(t, d.Cap(), 8) // Internal
}

func TestGenerics(t *testing.T) {
	a := MakeDequeSized[string](2)
	a.PushBack("test-a1")
	a.PushBack("test-a2")
	expect(t, a.PopBack(), "test-a2")
	expect(t, a.PopBack(), "test-a1")
	b := MakeDequeSized[float64](2)
	b.PushBack(3.1415927)
	b.PushBack(2.71828)
	expect(t, b.PopBack(), 2.71828)   // warn: float-cmp
	expect(t, b.PopBack(), 3.1415927) // warn: float-cmp
	c := MakeDequeSized[byte](2)
	c.PushBack(byte(13))
	c.PushBack(byte(10))
	expect(t, c.PopBack(), byte(10))
	expect(t, c.PopBack(), byte(13))
}
