package vessels

import (
	"testing"
)

func TestNewList(t *testing.T) {
	l := NewList[int]()
	expect(t, l.Len(), 0)
}

func TestNewListNode(t *testing.T) {
	n := NewListNode(42)
	expect(t, n.value, 42)
	expect(t, n.next, n)
	expect(t, n.prev, n)
}

func TestListNode(t *testing.T) {
	n := NewListNode(42)
	expect(t, n.next, n)
	expect(t, n.prev, n)
	// Next/Prev work on valid pointer to valid nodes
	l := NewListNode(0)
	r := NewListNode(0)
	n = NewListNode(0)
	l.next = n
	n.next = r
	r.prev = n
	n.prev = l

	expect(t, n.Next(), r)
	expect(t, n.Prev(), l)
}

func TestListNode_internal_insertBefore(t *testing.T) {
	n := NewListNode(9)
	a := NewListNode(7)
	b := NewListNode(8)
	a.insertBefore(n)
	b.insertBefore(n)
	expect(t, n.prev.value, 8)
	expect(t, n.prev.prev.value, 7)
}

func TestListNode_internal_insertAfter(t *testing.T) {
	n := NewListNode(9)
	a := NewListNode(7)
	b := NewListNode(8)
	a.insertAfter(n)
	b.insertAfter(n)
	expect(t, n.next.value, 8)
	expect(t, n.next.next.value, 7)
}

func TestListNode_internal_remove(t *testing.T) {
	n := NewListNode(8)
	a := NewListNode(7)
	b := NewListNode(9)
	a.insertBefore(n)
	b.insertAfter(n)
	expect(t, n.remove().value, 8)
	expect(t, a.next, b)
	expect(t, b.prev, a)
}

func TestListNode_internal_Swap(t *testing.T) {
	a := NewListNode(3)
	b := NewListNode(9)
	b.insertAfter(a)
	expect(t, a.next, b)
	expect(t, b.prev, a)
	expect(t, a.value, 3)
	expect(t, b.value, 9)
	a.Swap(b)
	expect(t, a.value, 9)
	expect(t, b.value, 3)
	expect(t, a.next, b)
	expect(t, b.prev, a)
}

func TestListLen(t *testing.T) {
	l := NewList[int]()
	expect(t, l.Len(), 0)
	l.PushBack(9)
	l.PushBack(8)
	l.PushBack(7)
	expect(t, l.Len(), 3)
}

func TestListIsEmpty(t *testing.T) {
	l := NewList[int]()
	expect(t, l.isEmpty(), true)
	l.PushBack(9)
	expect(t, l.isEmpty(), false)
}

func TestListSwap(t *testing.T) {
	a := NewList[int]()
	a.PushBack(9)
	a.PushBack(8)
	a.PushBack(7)
	b := NewList[int]()
	b.PushBack(6)
	b.PushBack(5)
	b.PushBack(4)

	a.Swap(b)
	expect(t, a.Len(), 3)
	expect(t, a.At(0), 6)
	expect(t, a.At(1), 5)
	expect(t, a.At(2), 4)
}

func TestListFront(t *testing.T) {
	l := NewList[int]()
	l.PushBack(9)
	expect(t, l.Front(), 9)
	l.PushFront(8)
	expect(t, l.Front(), 8)

	defer func() { _ = recover() }()
	l = NewList[int]()
	l.Front()
	t.Fatal("list Front() should panic on an empty list")
}

func TestListBack(t *testing.T) {
	l := NewList[int]()
	l.PushBack(9)
	expect(t, l.Back(), 9)
	l.PushFront(8)
	expect(t, l.Back(), 9)

	defer func() { _ = recover() }()
	l = NewList[int]()
	l.Back()
	t.Fatal("list Front() should panic on an empty list")
}

func TestListInsertBefore(t *testing.T) {
	l := NewList[int]()
	l.InsertBefore(8, l.End())
	expect(t, l.Len(), 1)
	l.InsertBefore(7, l.Begin())
	l.InsertBefore(9, l.End())
	expect(t, l.Len(), 3)
	expect(t, l.At(0), 7)
	expect(t, l.At(1), 8)
	expect(t, l.At(2), 9)
}

func TestListInsertAfter(t *testing.T) {
	l := NewList[int]()
	l.InsertAfter(9, l.head)
	expect(t, l.Len(), 1)
	l.InsertAfter(7, l.End())
	l.InsertAfter(8, l.Begin())
	expect(t, l.Len(), 3)
	expect(t, l.At(0), 7)
	expect(t, l.At(1), 8)
	expect(t, l.At(2), 9)
}

func TestListPushBack(t *testing.T) {
	l := NewList[int]()
	l.PushBack(9)
	expect(t, l.Len(), 1)
	l.PushBack(8)
	l.PushBack(7)
	expect(t, l.Len(), 3)

	expect(t, l.At(0), 9)
	expect(t, l.At(1), 8)
	expect(t, l.At(2), 7)
}

func TestListPopBack(t *testing.T) {
	l := NewList[int]()
	l.PushBack(9)
	l.PushBack(8)
	l.PushBack(7)

	expect(t, l.PopBack(), 7)
	expect(t, l.PopBack(), 8)
	expect(t, l.PopBack(), 9)

	defer func() { _ = recover() }()
	l.PopBack()
	t.Fatal("List PopBack() should panic when empty")
}

func TestListPushFront(t *testing.T) {
	l := NewList[int]()
	l.PushFront(9)
	l.PushFront(8)
	l.PushFront(7)
	expect(t, l.Len(), 3)

	expect(t, l.PopBack(), 9)
	expect(t, l.PopBack(), 8)
	expect(t, l.PopBack(), 7)
	expect(t, l.Len(), 0)
}

func TestListPopFront(t *testing.T) {
	l := NewList[int]()
	l.PushBack(9)
	l.PushBack(8)
	l.PushBack(7)

	expect(t, l.PopFront(), 9)
	expect(t, l.PopFront(), 8)
	expect(t, l.PopFront(), 7)

	defer func() { _ = recover() }()
	l.PopFront()
	t.Fatal("List PopFront() should panic when empty")
}

func TestListAppend(t *testing.T) {
	l := NewList[int]()
	l.Append(9)
	l.Append(8)
	l.Append(7)
	expect(t, l.Len(), 3)
	expect(t, l.At(0), 9)
	expect(t, l.At(1), 8)
	expect(t, l.At(2), 7)
}

func TestListClear(t *testing.T) {
	l := NewList[int]()
	l.PushBack(9)
	l.PushBack(8)
	l.PushBack(7)
	l.Clear()
	expect(t, l.Len(), 0)
	expect(t, l.Begin(), l.head)
	expect(t, l.End(), l.head)
}

func TestListAt(t *testing.T) {
	l := NewList[int]()
	l.PushBack(9)
	l.PushBack(8)
	l.PushBack(7)
	expect(t, l.At(0), 9)
	expect(t, l.At(1), 8)
	expect(t, l.At(2), 7)

	defer func() { _ = recover() }()
	l.At(l.Len())
	t.Fatal("List At() should panic with invalid index")
}

func TestListReverse(t *testing.T) {
	l := NewList[int]()
	l.PushBack(9)
	l.PushBack(8)
	l.PushBack(7)
	l.PushBack(6)
	l.PushBack(5)
	expect(t, l.At(0), 9)
	expect(t, l.At(1), 8)
	expect(t, l.At(2), 7)
	expect(t, l.At(3), 6)
	expect(t, l.At(4), 5)
	l.Reverse()
	expect(t, l.At(0), 5)
	expect(t, l.At(1), 6)
	expect(t, l.At(2), 7)
	expect(t, l.At(3), 8)
	expect(t, l.At(4), 9)
	l.Clear()
	l.PushBack(1)
	l.Reverse()
	expect(t, l.At(0), 1)
}
