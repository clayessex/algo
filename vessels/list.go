package vessels

import "cmp"

// List node holds a single value and pointers to the next and prev nodes
// The list head is a node where:
//
//	next = first node in the list
//	prev = last node in the list
//	next = prev = head when the list is empty
type ListNode[T any] struct {
	next  *ListNode[T]
	prev  *ListNode[T]
	value T
}

// Create a new list node, optionally with a value
func NewListNode[T any](v ...T) *ListNode[T] {
	n := &ListNode[T]{}
	n.next = n
	n.prev = n
	if len(v) >= 1 {
		n.value = v[0]
	}
	return n
}

// Pointer to the next node in the list
func (n *ListNode[T]) Next() *ListNode[T] {
	return n.next
}

// Pointer to the previous node in the list
func (n *ListNode[T]) Prev() *ListNode[T] {
	return n.prev
}

// insert n before p
func (n *ListNode[T]) insertBefore(p *ListNode[T]) *ListNode[T] {
	p.prev.next = n
	n.prev = p.prev
	n.next = p
	p.prev = n
	return n
}

// insert n after p
func (n *ListNode[T]) insertAfter(p *ListNode[T]) *ListNode[T] {
	p.next.prev = n
	n.next = p.next
	n.prev = p
	p.next = n
	return n
}

// remove n from the list
// nil the pointers for GC (shouldn't be needed)
func (n *ListNode[T]) remove() {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.prev = nil
	n.next = nil
}

// Moves [first, last) before pos
func splice[T any](pos, first, last *ListNode[T]) {
	if pos == first || pos == last {
		return
	}
	oleft := first.prev
	oright := last
	last = last.prev
	left := pos.prev
	right := pos

	// cut from old
	oleft.next = oright
	oright.prev = oleft

	// splice to new
	left.next = first
	first.prev = left
	right.prev = last
	last.next = right
}

// Swap the two nodes without touching their values
func (n *ListNode[T]) Swap(o *ListNode[T]) {
	tmp := o.next
	splice(n, o, o.next)
	if tmp != n {
		splice(tmp, n, n.next)
	}
}

// The list head is a node where:
//
//	next = first node in the list
//	prev = last node in the list
//	next = prev = head when the list is empty
type List[T any] struct {
	head *ListNode[T]
	len  int
}

// Create a new list
func NewList[T any]() *List[T] {
	list := List[T]{}
	list.head = NewListNode[T]()
	return &list
}

// Length of the list
func (list *List[T]) Len() int {
	return list.len
}

// Swap two lists
func (list *List[T]) Swap(o *List[T]) {
	*list, *o = *o, *list
}

// Len() == 0
func (list *List[T]) isEmpty() bool {
	return list.len == 0
}

// First element of the list
func (list *List[T]) Begin() *ListNode[T] {
	return list.head.next
}

// Node following the last in the list
func (list *List[T]) End() *ListNode[T] {
	return list.head
}

// First value of the list
func (list *List[T]) Front() T {
	if list.len == 0 {
		panic("list Front() called on empty list")
	}
	return list.head.next.value
}

// Last value of the list
func (list *List[T]) Back() T {
	if list.len == 0 {
		panic("list Back() called on empty list")
	}
	return list.head.prev.value
}

// Insert the value into the list before node pos
func (list *List[T]) insert(v T, pos *ListNode[T]) *ListNode[T] {
	n := NewListNode(v)
	list.len++
	return n.insertBefore(pos)
}

// Remove the node from the list
func (list *List[T]) RemoveNode(pos *ListNode[T]) {
	list.len--
	pos.remove()
}

// Insert the value into the list before pos
func (list *List[T]) InsertBefore(v T, pos *ListNode[T]) *ListNode[T] {
	return list.insert(v, pos)
}

// Insert the value into the list after pos
func (list *List[T]) InsertAfter(v T, pos *ListNode[T]) *ListNode[T] {
	return list.insert(v, pos.Next())
}

// Add a new value onto the end of the list
func (list *List[T]) PushBack(v T) *ListNode[T] {
	return list.insert(v, list.End())
}

// Add a new value onto the beginning of the list
func (list *List[T]) PushFront(v T) *ListNode[T] {
	return list.insert(v, list.Begin())
}

// Remove the last value from the list and return it
func (list *List[T]) PopBack() T {
	if list.len == 0 {
		panic("list PopBack() called on empty list")
	}
	p := list.End().prev
	value := p.value
	list.RemoveNode(p)
	return value
}

// Remove the first value from the list and return it
func (list *List[T]) PopFront() T {
	if list.len == 0 {
		panic("list PopFront() called on empty list")
	}
	value := list.Begin().value
	list.RemoveNode(list.Begin())
	return value
}

// Append one or more values to the list
func (list *List[T]) Append(v ...T) {
	for _, el := range v {
		list.PushBack(el)
	}
}

// Return a slice containing the list values
func (list *List[T]) Values() []T {
	result := make([]T, 0, list.Len())
	if list.Len() == 0 {
		return result
	}
	first := list.Begin()
	for first != list.End() {
		result = append(result, first.value)
		first = first.Next()
	}
	return result
}

// Remove all the values from the list
func (list *List[T]) Clear() {
	for p := list.Begin(); p != list.End(); {
		r := p
		p = p.Next()
		r.remove()
	}
	list.len = 0
}

// Return the value at index offset into the list
// The list is not internally indexed so this function has O(n) complexity
func (list *List[T]) At(index int) T {
	if index < 0 || index >= list.len {
		panic("list At called with index out of bounds")
	}
	p := list.Begin()
	for ; index > 0; index-- {
		p = p.Next()
	}

	return p.value
}

// Reverse the elements of the list
// Requires n/2 swaps
func (list *List[T]) Reverse() {
	if list.Len() <= 1 {
		return
	}

	front := list.head.next
	back := list.head.prev

	// Swap the nodes from the front of the list with the ones at the back
	for i := 0; i < list.Len()/2; i++ {
		a, b := front, back
		front, back = front.next, back.prev
		a.Swap(b)
	}
}

// Remove the sequence [first, last) from srcList and insert it into list before pos
func (list *List[T]) Splice(pos *ListNode[T], srcList *List[T], first, last *ListNode[T]) {
	count := 0
	for p := first; p != last; p = p.next {
		count++
	}
	splice(pos, first, last)
	srcList.len -= count
	list.len += count
}

// Remove values from the list where pred(value) is true
func ListRemoveFunc[T any](list *List[T], pred func(v T) bool) int {
	if list.Len() == 0 {
		return 0
	}

	count := 0
	p := list.Begin()
	for p != list.End() {
		if pred(p.value) {
			tmp := p
			p = p.Next()
			tmp.remove()
			list.len--
			count++
		} else {
			p = p.Next()
		}
	}

	return count
}

// Remove values from the list that match value using ==
// Requires T to be comparable which is not a requirement of the underlying list
func ListRemove[T comparable](list *List[T], value T) int {
	return ListRemoveFunc(list, func(v T) bool {
		return v == value
	})
}

// Remove consecutive values from the list where pred(a, b) is true
func ListUniqueFunc[T any](list *List[T], pred func(a, b T) bool) int {
	if list.Len() <= 1 {
		return 0
	}

	count := 0
	first := list.Begin().Next()
	for first != list.End() {
		if pred(first.Prev().value, first.value) {
			tmp := first
			first = first.Next()
			tmp.remove()
			list.len--
			count++
		} else {
			first = first.Next()
		}
	}

	return count
}

// Remove consecutive values from the list using ==
// Requires T to be comparable which is not a requirement of the underlying list
func ListUnique[T comparable](list *List[T]) int {
	return ListUniqueFunc(list, func(a, b T) bool {
		return a == b
	})
}

// Merge sorted list "other" into sorted list "list" where comp(a, b) is true
func ListMergeFunc[T any](list, other *List[T], comp func(a, b T) bool) {
	if list == other || other.Len() == 0 {
		return
	}

	first1 := list.Begin()
	last1 := list.End()
	first2 := other.Begin()
	last2 := other.End()

	for first1 != last1 && first2 != last2 {
		if comp(first2.value, first1.value) {
			next := first2.next
			for next != last2 {
				if !comp(next.value, first1.value) {
					break
				}
				next = next.next
			}
			splice(first1, first2, next)
			first2 = next
		} else {
			first1 = first1.next
		}
	}

	if first2 != last2 {
		splice(last1, first2, last2)
	}
	list.len += other.len
	other.len = 0
}

// Merge sorted list "other" into sorted list "list" using cmp.Less
// Requires that T be cmp.Ordered, which is not a requirement of the underlying list
func ListMerge[T cmp.Ordered](list, other *List[T]) {
	ListMergeFunc(list, other, cmp.Less)
}

// Find a value in list where pred(value) is true
// Returns the (node, ok) if found, otherwise it returns
// (list.End(), false)
func ListFindFunc[T any](list *List[T], pred func(a T) bool) (*ListNode[T], bool) {
	first := list.Begin()
	for first != list.End() {
		if pred(first.value) {
			return first, true
		}
		first = first.next
	}
	return list.End(), false
}

// Find a value v in list using ==
// Returns the (node, ok) if found, otherwise it returns
// (list.End(), false)
func ListFind[T comparable](list *List[T], v T) (*ListNode[T], bool) {
	return ListFindFunc(list, func(a T) bool {
		return a == v
	})
}

// Run a function f on every value in the list
func (list *List[T]) ForEach(f func(v T)) {
	for p := list.Begin(); p != list.End(); p = p.Next() {
		f(p.value)
	}
}
