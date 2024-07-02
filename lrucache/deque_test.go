package main

import (
	"reflect"
	"testing"
)

func dequeBufferContains(d *Deque, want []int) bool {
	index := d.tail
	for _, v := range want {
		if v != d.buf[index] {
			return false
		}
		index = (index + 1) % d.cap
	}
	return true
}

func makeTestDeque(values []int, head int, tail int) *Deque {
	return &Deque{values, len(values), head, tail}
}

func expect(t *testing.T, got interface{}, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("%s failed: got: %v != want: %v\n", t.Name(), got, want)
	}
}

func TestMakeDequeSized(t *testing.T) {
	d := MakeDequeSized(4)
	if d.cap != 4 {
		t.Fatalf("MakeDequeSized failed: expected cap: %v, got %v", 4, d.cap)
	}
	if len(d.buf) != 4 {
		t.Fatalf("MakeDequeSized failed: expected len: %v, got %v", 4, len(d.buf))
	}
}

func TestDequeSize(t *testing.T) {
	d := &Deque{[]int{0, 0}, 2, 0, 0}
	expect(t, d.Size(), 0)
	d = &Deque{[]int{9, 0}, 2, 1, 0}
	expect(t, d.Size(), 1)
	d = &Deque{[]int{9, 8, 0, 0}, 4, 2, 0}
	expect(t, d.Size(), 2)
	d = &Deque{[]int{0, 9, 8, 0}, 4, 3, 1}
	expect(t, d.Size(), 2)
	d = &Deque{[]int{8, 7, 0, 9}, 4, 2, 3}
	expect(t, d.Size(), 3)
}

func TestPushBack(t *testing.T) {
	d := &Deque{[]int{0, 0}, 2, 0, 0}
	d.PushBack(9)
	expect(t, d, &Deque{[]int{9, 0}, 2, 1, 0})
	d.PushBack(8)
	expect(t, d, &Deque{[]int{9, 8, 0, 0}, 4, 2, 0})
	d.PushBack(7)
	expect(t, d, &Deque{[]int{9, 8, 7, 0}, 4, 3, 0})

	d = &Deque{[]int{0, 8, 7, 0}, 4, 3, 1}
	d.PushBack(6)
	expect(t, d, &Deque{[]int{0, 8, 7, 6}, 4, 0, 1})
	d.PushBack(5)
	expect(t, d, &Deque{[]int{8, 7, 6, 5, 0, 0, 0, 0}, 8, 4, 0})

	d = &Deque{[]int{7, 0, 0, 8}, 4, 1, 3}
	d.PushBack(6)
	expect(t, d, &Deque{[]int{7, 6, 0, 8}, 4, 2, 3})
}

// func TestPushBackEmpty(t *testing.T) {
// 	d := MakeDequeSized(5)
// 	d.PushBack(9)
//
// 	target := &Deque{
// 		[]int{9, 0, 0, 0, 0},
// 		5,
// 		1,
// 		0,
// 	}
//
// 	if !reflect.DeepEqual(d, target) {
// 		t.Fatalf("Deque PushBackEmpty failed: %v, %v\n", d, target)
// 	}
// }
//
// func TestPushBackFull(t *testing.T) {
// 	d := MakeDequeSized(2)
// 	d.PushBack(9)
// 	d.PushBack(9)
// 	d.PushBack(9)
//
// 	target := &Deque{
// 		[]int{9, 9, 9, 0},
// 		4,
// 		3,
// 		0,
// 	}
//
// 	if !reflect.DeepEqual(d, target) {
// 		t.Fatalf("Deque PushBackFull failed: %v, %v\n", d, target)
// 	}
// }
