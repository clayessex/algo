package vessels

import (
	"testing"
)

func TestNewList(t *testing.T) {
	l := NewList[int]()
	expect(t, l.Len(), 0)
}

func TestListNode(t *testing.T) {
	// Next/Prev work on nil pointer
	var n *ListNode[int] = nil
	expectNil(t, n.Next())
	expectNil(t, n.Prev())

	// Next/Prev work on valid pointer to nil nodes
	n = &ListNode[int]{nil, nil, 0}
	expectNil(t, n.Next())
	expectNil(t, n.Prev())

	// Next/Prev work on valid pointer to valid nodes
	l := &ListNode[int]{nil, nil, 0}
	r := &ListNode[int]{nil, nil, 0}
	n = &ListNode[int]{nil, nil, 0}
	l.next = n
	n.next = r
	r.prev = n
	n.prev = l

	expect(t, n.Next(), r)
	expect(t, n.Prev(), l)
}

func TestListLen(t *testing.T) {
	l := NewList[int]()
	expect(t, l.Len(), 0)
	l.PushBack(9)
	l.PushBack(8)
	l.PushBack(7)
	expect(t, l.Len(), 3)
}

func TestListInsertBefore(t *testing.T) {
	l := NewList[int]()
	l.InsertBefore(9, nil)
	expect(t, l.Len(), 1)
	l.InsertBefore(7, l.Front())
	l.InsertBefore(8, l.Back())
	expect(t, l.Len(), 3)
	expect(t, l.At(0), 7)
	expect(t, l.At(1), 8)
	expect(t, l.At(2), 9)
}

func TestListInsertAfter(t *testing.T) {
	l := NewList[int]()
	l.InsertAfter(9, nil)
	expect(t, l.Len(), 1)
	l.InsertAfter(8, l.Front())
	l.InsertAfter(7, l.Back())
	expect(t, l.Len(), 3)
	expect(t, l.At(0), 9)
	expect(t, l.At(1), 8)
	expect(t, l.At(2), 7)
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
	expectNil(t, l.Front())
	expectNil(t, l.Back())
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
