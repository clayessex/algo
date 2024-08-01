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
	n.remove()
	expect(t, a.next, b)
	expect(t, b.prev, a)
}

func TestListNode_Swap(t *testing.T) {
	list := NewList[int]()
	list.Append(3, 9, 7)
	mid := list.Begin().Next()

	expect(t, list.Begin().Next(), mid)
	expect(t, list.End().Prev().Prev(), mid)

	expect(t, list.Begin().value, 3)
	expect(t, list.End().Prev().value, 7)

	list.Begin().Swap(list.End().Prev())

	// ensure that the values swapped
	expect(t, list.Begin().value, 7)
	expect(t, list.End().Prev().value, 3)

	// ensure that the node positions swapped
	expect(t, list.Begin().Next(), mid)
	expect(t, list.End().Prev().Prev(), mid)
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
	expect(t, l.Len(), 1)
	expect(t, l.At(0), 9)
	l.Append(8, 7, 6)
	expect(t, l.Len(), 4)
	expect(t, l.At(0), 9)
	expect(t, l.At(1), 8)
	expect(t, l.At(2), 7)
	expect(t, l.At(3), 6)
}

func TestListValues(t *testing.T) {
	list := NewList[int]()
	sl := list.Values()
	expect(t, len(sl), 0)
	list.Append(9, 8, 7, 6, 1)
	sl = list.Values()
	expect(t, sl, []int{9, 8, 7, 6, 1})
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
	expect(t, l.Len(), 5)
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
	l.PushBack(2)
	l.Reverse()
	expect(t, l.At(0), 2)
	expect(t, l.At(1), 1)
}

func TestList_internal_splice(t *testing.T) {
	a := NewList[int]()
	a.PushBack(9)
	a.PushBack(8)
	a.PushBack(7)
	b := NewList[int]()
	b.PushBack(6)
	b.PushBack(5)
	b.PushBack(4)

	splice(a.Begin().Next(), b.Begin(), b.End())
	a.len += 3

	expect(t, a.Len(), 6)

	expect(t, a.At(0), 9)
	expect(t, a.At(1), 6)
	expect(t, a.At(2), 5)
	expect(t, a.At(3), 4)
	expect(t, a.At(4), 8)
	expect(t, a.At(5), 7)

	splice(a.Begin(), a.Begin(), a.Begin().Next())
	expect(t, a.Len(), 6)
	expect(t, len(a.Values()), 6)
}

func TestListSplice(t *testing.T) {
	list := NewList[int]()
	list.Append(1, 2, 3)
	list.Splice(list.Begin(), list, list.Begin().Next(), list.End().Prev())
	expect(t, list.Values(), []int{2, 1, 3})
	expect(t, list.Len(), 3)
	list.Clear()
	list2 := NewList[int]()

	list.Append(4, 5, 6)
	list2.Append(1, 2, 3)
	list.Splice(list.Begin(), list2, list2.Begin(), list2.End())
	expect(t, list.Values(), []int{1, 2, 3, 4, 5, 6})
	expect(t, list.Len(), 6)
	expect(t, list2.Len(), 0)
}

func TestListRemove(t *testing.T) {
	list := NewList[int]()
	ListRemove(list, 99)
	list.Append(4, 5, 6)
	count := ListRemove(list, 5)
	expect(t, list.At(0), 4)
	expect(t, list.At(1), 6)
	expect(t, count, 1)
}

func TestListUnique(t *testing.T) {
	list := NewList[int]()
	count := ListUnique(list)
	expect(t, count, 0)
	list.Append(4, 4, 5, 6, 6, 7)
	count = ListUnique(list)
	expect(t, count, 2)
	expect(t, list.Len(), 4)
	expect(t, list.Values(), []int{4, 5, 6, 7})
}

func TestListMerge(t *testing.T) {
	a := NewList[int]()
	b := NewList[int]()
	a.Append(1, 4, 5, 7)
	b.Append(2, 3, 6, 8)
	ListMerge(a, b)
	expect(t, a.Len(), 8)
	expect(t, b.Len(), 0)
	expect(t, a.Values(), []int{1, 2, 3, 4, 5, 6, 7, 8})

	a.Clear()
	b.Clear()
	a.Append(1, 2, 3)
	b.Append(6, 7, 8)
	ListMerge(a, b)
	expect(t, a.Len(), 6)
	expect(t, a.Values(), []int{1, 2, 3, 6, 7, 8})

	a.Clear()
	b.Clear()
	b.Append(6, 7, 8)
	ListMerge(a, b)
	expect(t, a.Len(), 3)
	expect(t, b.Len(), 0)
	expect(t, a.Values(), []int{6, 7, 8})

	a.Clear()
	b.Clear()
	a.Append(6, 7, 8)
	ListMerge(a, b)
	expect(t, a.Len(), 3)
	expect(t, b.Len(), 0)
	expect(t, a.Values(), []int{6, 7, 8})
}

func TestListFind(t *testing.T) {
	list := NewList[int]()
	list.Append(1, 2, 3, 4, 5)
	node, ok := ListFind(list, 2)
	expect(t, ok, true)
	expect(t, node, list.Begin().Next())
	node, ok = ListFind(list, 9)
	expect(t, ok, false)
	expect(t, node, list.End())
}

func TestListRange(t *testing.T) {
	list := NewList[int]()
	list.Append(1, 2, 3, 4)
	sum := 0
	list.Range(func(v int) {
		sum += v
	})
	expect(t, sum, 10)
}
