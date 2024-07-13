package vessels

import (
	"cmp"
	"fmt"
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
}

func advance(n *ListNode[int], sz int) *ListNode[int] {
	for i := 0; i < sz; i++ {
		n = n.next
	}
	return n
}

func TestList_internal_mergeNodes(t *testing.T) {
	tests := []struct {
		name        string
		input, want []int
	}{
		{"a", []int{4, 7, 8, 5, 6, 9}, []int{4, 5, 6, 7, 8, 9}},
		{"b", []int{7, 8, 9, 4, 5, 6}, []int{4, 5, 6, 7, 8, 9}},
		{"c", []int{4, 5, 6, 7, 8, 9}, []int{4, 5, 6, 7, 8, 9}},
		{"d", []int{4, 8, 9, 5, 6, 7}, []int{4, 5, 6, 7, 8, 9}},
		{"e", []int{2, 4, 9, 1, 5, 7}, []int{1, 2, 4, 5, 7, 9}},
		{"f", []int{7, 8, 9, 1, 2, 3}, []int{1, 2, 3, 7, 8, 9}},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			list := NewList[int]()
			list.Append(v.input...)

			mid := advance(list.Begin(), list.Len()/2)
			end := mergeOrderedNodes(list.Begin(), mid, list.End())

			expect(t, end, list.Begin())
			compareNodesSlice(t, list.Begin(), list.End(), v.want)
		})
	}
}

func TestList_internal_mergeNodesPartail(t *testing.T) {
	input, want := []int{4, 7, 3, 6, 9, 8}, []int{3, 4, 6, 7, 9, 8}

	list := NewList[int]()
	list.Append(input...)

	mid := advance(list.Begin(), 2)
	last := advance(mid, 2)

	// merge first two pairs
	_ = mergeOrderedNodes(list.Begin(), mid, last)

	logListNodes(t, list.Begin(), list.End())
	compareNodesSlice(t, list.Begin(), list.End(), want)
}

func TestList_internal_sortNodesEmpty(t *testing.T) {
	list := NewList[int]()
	begin, end := sortNodes(list.End(), 0)
	expect(t, begin, list.End())
	expect(t, end, list.End())
}

func TestListSortList(t *testing.T) {
	tests := []struct {
		name string
		data []int
	}{
		{"a", []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{"b", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"c", []int{8, 2, 6, 5, 1, 4, 3, 9, 7, 0}},
		{"d", []int{}},
		{"e", []int{8}},
		{"f", []int{9, 8}},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			list := NewList[int]()
			list.Append(v.data...)
			SortList(list)
			expectOrdered(t, list.Begin(), list.End())
		})
	}
}

func isOrdered[T cmp.Ordered](first *ListNode[T], last *ListNode[T]) bool {
	if first == last {
		return true
	}
	first = first.next
	for first != last {
		if !(first.prev.value <= first.value) {
			return false
		}
		first = first.next
	}
	return true
}

func intsToBytes(values []int) []byte {
	result := make([]byte, len(values))
	for i := 0; i < len(values); i++ {
		result[i] = byte(values[i])
	}
	return result
}

func printBytes(values []byte) string {
	result := "["
	for i := 0; i < min(len(values), 100); i++ {
		result = result + fmt.Sprintf("%v ", int(values[i]))
	}
	if len(values) > 100 {
		result = result + "..."
	}
	result += "]"
	return result
}

func FuzzListSortList(f *testing.F) {
	f.Add([]byte{})
	f.Add([]byte{1})
	f.Add([]byte{2, 1})
	f.Add([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9})
	f.Add([]byte{9, 8, 7, 6, 5, 4, 3, 2, 1})
	f.Add([]byte{
		249, 139, 93, 55, 25, 158, 198, 77, 67, 68, 255, 128, 135, 147, 198, 219,
		59, 243, 210, 70, 112, 15, 68, 69, 39, 166, 132, 161, 233, 83, 166, 126,
		74, 11, 129, 186, 26, 29, 149, 150, 247, 172, 99, 235, 12, 128, 131, 183,
		82, 56, 221, 162, 211, 216, 91, 157, 214, 29, 195, 97, 189, 250, 205, 85,
		78, 185, 60, 29, 160, 49, 16, 32, 253, 211, 141, 162, 251, 183, 200, 151,
		142, 197, 197, 88, 49, 3, 210, 197, 92, 9, 251, 222, 108, 60, 234, 34,
		25, 6, 142, 238,
	})

	f.Fuzz(func(t *testing.T, data []byte) {
		list := NewList[byte]()
		for _, v := range data {
			list.Append(v)
		}

		SortList(list)

		if !isOrdered(list.Begin(), list.End()) {
			t.Log(data)
			t.Fatal("sorted list is not sorted")
		}
	})
}
