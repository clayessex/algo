package vessels

import (
	"cmp"
	"testing"
)

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
			end := mergeOrderedNodes(list.Begin(), mid, list.End(), cmp.Less)

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
	_ = mergeOrderedNodes(list.Begin(), mid, last, cmp.Less)

	logListNodes(t, list.Begin(), list.End())
	compareNodesSlice(t, list.Begin(), list.End(), want)
}

func TestList_internal_sortNodesEmpty(t *testing.T) {
	list := NewList[int]()
	begin, end := sortNodes(list.End(), 0, cmp.Less)
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

func TestListSortListFunc(t *testing.T) {
	list := NewList[int]()
	list.Append([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	SortListFunc(list, func(a, b int) bool {
		return b < a
	})
	if !isReverseOrdered(list.Begin(), list.End()) {
		t.Fatal("list is not reverse ordered")
	}
}

func TestListSortListStrings(t *testing.T) {
	list := NewList[string]()
	list.Append([]string{"banana", "apple", "foo", "bar", "pear", "strawberry"}...)
	SortList(list)
	if !isOrdered(list.Begin(), list.End()) {
		t.Fatal("string list is not ordered")
	}
}

func isReverseOrdered[T cmp.Ordered](first, last *ListNode[T]) bool {
	if first == last {
		return true
	}
	first = first.next
	for first != last {
		if !(first.prev.value > first.value) {
			return false
		}
		first = first.next
	}
	return true
}

func isOrdered[T cmp.Ordered](first, last *ListNode[T]) bool {
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
