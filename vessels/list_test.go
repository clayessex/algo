package vessels

import (
	"testing"

	"github.com/clayessex/algo/expected"
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
	x := expected.New(t)
	a := NewList[int]()
	a.PushBack(9)
	a.PushBack(8)
	a.PushBack(7)
	b := NewList[int]()
	b.PushBack(6)
	b.PushBack(5)
	b.PushBack(4)

	a.Swap(b)
	x.Expect(a.Len()).ToBe(3)
	x.ExpectOk(a.At(0)).ToBe(6)
	x.ExpectOk(a.At(1)).ToBe(5)
	x.ExpectOk(a.At(2)).ToBe(4)
}

func TestListFront(t *testing.T) {
	x := expected.New(t)
	l := NewList[int]()
	l.PushBack(9)
	x.ExpectOk(l.Front()).ToBe(9)
	l.PushFront(8)
	x.ExpectOk(l.Front()).ToBe(8)

	l = NewList[int]()
	x.ExpectNotOk(l.Front())
}

func TestListBack(t *testing.T) {
	x := expected.New(t)
	l := NewList[int]()
	l.PushBack(9)
	x.ExpectOk(l.Back()).ToBe(9)
	l.PushFront(8)
	x.ExpectOk(l.Back()).ToBe(9)

	l = NewList[int]()
	x.ExpectNotOk(l.Back())
}

func TestListInsertBefore(t *testing.T) {
	x := expected.New(t)
	l := NewList[int]()
	l.InsertBefore(8, l.End())
	x.Expect(l.Len()).ToBe(1)
	l.InsertBefore(7, l.Begin())
	l.InsertBefore(9, l.End())
	x.Expect(l.Len()).ToBe(3)
	x.ExpectOk(l.At(0)).ToBe(7)
	x.ExpectOk(l.At(1)).ToBe(8)
	x.ExpectOk(l.At(2)).ToBe(9)
}

func TestListInsertAfter(t *testing.T) {
	x := expected.New(t)
	l := NewList[int]()
	l.InsertAfter(9, l.head)
	x.Expect(l.Len()).ToBe(1)
	l.InsertAfter(7, l.End())
	l.InsertAfter(8, l.Begin())
	x.Expect(l.Len()).ToBe(3)
	x.ExpectOk(l.At(0)).ToBe(7)
	x.ExpectOk(l.At(1)).ToBe(8)
	x.ExpectOk(l.At(2)).ToBe(9)
}

func TestListPushBack(t *testing.T) {
	x := expected.New(t)
	l := NewList[int]()
	l.PushBack(9)
	x.Expect(l.Len()).ToBe(1)
	l.PushBack(8)
	l.PushBack(7)
	x.Expect(l.Len()).ToBe(3)

	x.ExpectOk(l.At(0)).ToBe(9)
	x.ExpectOk(l.At(1)).ToBe(8)
	x.ExpectOk(l.At(2)).ToBe(7)
}

func TestListPopBack(t *testing.T) {
	x := expected.New(t)
	l := NewList[int]()
	l.PushBack(9)
	l.PushBack(8)
	l.PushBack(7)

	x.ExpectOk(l.PopBack()).ToBe(7)
	x.ExpectOk(l.PopBack()).ToBe(8)
	x.ExpectOk(l.PopBack()).ToBe(9)

	x.ExpectNotOk(l.PopBack())
}

func TestListPushFront(t *testing.T) {
	x := expected.New(t)
	l := NewList[int]()
	l.PushFront(9)
	l.PushFront(8)
	l.PushFront(7)
	x.Expect(l.Len()).ToBe(3)

	x.ExpectOk(l.PopBack()).ToBe(9)
	x.ExpectOk(l.PopBack()).ToBe(8)
	x.ExpectOk(l.PopBack()).ToBe(7)
	x.Expect(l.Len()).ToBe(0)
}

func TestListPopFront(t *testing.T) {
	x := expected.New(t)
	l := NewList[int]()
	l.PushBack(9)
	l.PushBack(8)
	l.PushBack(7)

	x.ExpectOk(l.PopFront()).ToBe(9)
	x.ExpectOk(l.PopFront()).ToBe(8)
	x.ExpectOk(l.PopFront()).ToBe(7)

	x.ExpectNotOk(l.PopFront())
}

func TestListAppend(t *testing.T) {
	x := expected.New(t)
	l := NewList[int]()
	l.Append(9)
	x.Expect(l.Len()).ToBe(1)
	x.ExpectOk(l.At(0)).ToBe(9)
	l.Append(8, 7, 6)
	x.Expect(l.Len()).ToBe(4)
	x.ExpectOk(l.At(0)).ToBe(9)
	x.ExpectOk(l.At(1)).ToBe(8)
	x.ExpectOk(l.At(2)).ToBe(7)
	x.ExpectOk(l.At(3)).ToBe(6)
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
	x := expected.New(t)
	l := NewList[int]()
	l.PushBack(9)
	l.PushBack(8)
	l.PushBack(7)
	x.ExpectOk(l.At(0)).ToBe(9)
	x.ExpectOk(l.At(1)).ToBe(8)
	x.ExpectOk(l.At(2)).ToBe(7)
	x.ExpectNotOk(l.At(l.Len()))
}

func TestListReverse(t *testing.T) {
	x := expected.New(t)
	l := NewList[int]()
	l.PushBack(9)
	l.PushBack(8)
	l.PushBack(7)
	l.PushBack(6)
	l.PushBack(5)
	x.Expect(l.Len()).ToBe(5)
	x.ExpectOk(l.At(0)).ToBe(9)
	x.ExpectOk(l.At(1)).ToBe(8)
	x.ExpectOk(l.At(2)).ToBe(7)
	x.ExpectOk(l.At(3)).ToBe(6)
	x.ExpectOk(l.At(4)).ToBe(5)
	l.Reverse()
	x.ExpectOk(l.At(0)).ToBe(5)
	x.ExpectOk(l.At(1)).ToBe(6)
	x.ExpectOk(l.At(2)).ToBe(7)
	x.ExpectOk(l.At(3)).ToBe(8)
	x.ExpectOk(l.At(4)).ToBe(9)
	l.Clear()
	l.PushBack(1)
	l.Reverse()
	x.ExpectOk(l.At(0)).ToBe(1)
	l.PushBack(2)
	l.Reverse()
	x.ExpectOk(l.At(0)).ToBe(2)
	x.ExpectOk(l.At(1)).ToBe(1)
}

func TestList_internal_splice(t *testing.T) {
	x := expected.New(t)
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

	x.Expect(a.Len()).ToBe(6)

	x.ExpectOk(a.At(0)).ToBe(9)
	x.ExpectOk(a.At(1)).ToBe(6)
	x.ExpectOk(a.At(2)).ToBe(5)
	x.ExpectOk(a.At(3)).ToBe(4)
	x.ExpectOk(a.At(4)).ToBe(8)
	x.ExpectOk(a.At(5)).ToBe(7)

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
	x := expected.New(t)
	list := NewList[int]()
	ListRemove(list, 99)
	list.Append(4, 5, 6)
	count := ListRemove(list, 5)
	x.ExpectOk(list.At(0)).ToBe(4)
	x.ExpectOk(list.At(1)).ToBe(6)
	x.Expect(count).ToBe(1)
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
