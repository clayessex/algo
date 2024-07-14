package vessels

import "cmp"

type ListNode[T any] struct {
	next  *ListNode[T]
	prev  *ListNode[T]
	value T
}

func NewListNode[T any](v ...T) *ListNode[T] {
	n := &ListNode[T]{}
	n.next = n
	n.prev = n
	if len(v) >= 1 {
		n.value = v[0]
	}
	return n
}

func (n *ListNode[T]) Next() *ListNode[T] {
	return n.next
}

func (n *ListNode[T]) Prev() *ListNode[T] {
	return n.prev
}

/** insert n before p */
func (n *ListNode[T]) insertBefore(p *ListNode[T]) *ListNode[T] {
	p.prev.next = n
	n.prev = p.prev
	n.next = p
	p.prev = n
	return n
}

/** insert n after p */
func (n *ListNode[T]) insertAfter(p *ListNode[T]) *ListNode[T] {
	p.next.prev = n
	n.next = p.next
	n.prev = p
	p.next = n
	return n
}

func (n *ListNode[T]) remove() {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.prev = nil
	n.next = nil
}

/** Moves [first, last) before pos */
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

func (n *ListNode[T]) Swap(o *ListNode[T]) {
	tmp := o.next
	splice(n, o, o.next)
	if tmp != n {
		splice(tmp, n, n.next)
	}
}

type List[T any] struct {
	head *ListNode[T]
	len  int
}

func NewList[T any]() *List[T] {
	list := List[T]{}
	list.head = NewListNode[T]()
	return &list
}

func (list *List[T]) Len() int {
	return list.len
}

func (list *List[T]) Swap(o *List[T]) {
	*list, *o = *o, *list
}

func (list *List[T]) isEmpty() bool {
	return list.len == 0
}

func (list *List[T]) Begin() *ListNode[T] {
	return list.head.next
}

func (list *List[T]) End() *ListNode[T] {
	return list.head
}

func (list *List[T]) Front() T {
	if list.len == 0 {
		panic("list Front() called on empty list")
	}
	return list.head.next.value
}

func (list *List[T]) Back() T {
	if list.len == 0 {
		panic("list Back() called on empty list")
	}
	return list.head.prev.value
}

func (list *List[T]) insert(v T, pos *ListNode[T]) *ListNode[T] {
	n := NewListNode(v)
	list.len++
	return n.insertBefore(pos)
}

func (list *List[T]) RemoveNode(pos *ListNode[T]) {
	list.len--
	pos.remove()
}

func (list *List[T]) InsertBefore(v T, pos *ListNode[T]) *ListNode[T] {
	return list.insert(v, pos)
}

func (list *List[T]) InsertAfter(v T, pos *ListNode[T]) *ListNode[T] {
	return list.insert(v, pos.Next())
}

func (list *List[T]) PushBack(v T) *ListNode[T] {
	return list.insert(v, list.End())
}

func (list *List[T]) PushFront(v T) *ListNode[T] {
	return list.insert(v, list.Begin())
}

func (list *List[T]) PopBack() T {
	if list.len == 0 {
		panic("list PopBack() called on empty list")
	}
	p := list.End().prev
	value := p.value
	list.RemoveNode(p)
	return value
}

func (list *List[T]) PopFront() T {
	if list.len == 0 {
		panic("list PopFront() called on empty list")
	}
	value := list.Begin().value
	list.RemoveNode(list.Begin())
	return value
}

func (list *List[T]) Append(v ...T) {
	for _, el := range v {
		list.PushBack(el)
	}
}

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

func (list *List[T]) Clear() {
	for p := list.Begin(); p != list.End(); {
		r := p
		p = p.Next()
		r.remove()
	}
	list.len = 0
}

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

func (list *List[T]) Reverse() {
	if list.Len() <= 1 {
		return
	}
	p := list.End()
	for i := 0; i < list.len-1; i++ {
		splice(p, list.Begin(), list.Begin().next)
		p = p.prev
	}
}

func (list *List[T]) Splice(pos, first, last *ListNode[T]) {
	splice(pos, first, last)
}

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

func ListRemove[T comparable](list *List[T], value T) int {
	return ListRemoveFunc(list, func(v T) bool {
		return v == value
	})
}

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

func ListUnique[T comparable](list *List[T]) int {
	return ListUniqueFunc(list, func(a, b T) bool {
		return a == b
	})
}

func ListMerge[T cmp.Ordered](list, other *List[T]) {
	if list == other || other.Len() == 0 {
		return
	}

	first1 := list.Begin()
	last1 := list.End()
	first2 := other.Begin()
	last2 := other.End()

	for first1 != last1 && first2 != last2 {
		if first2.value < first1.value {
			next := first2.next
			for next != last2 {
				if !(next.value < first1.value) {
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

// TODO: iterations
// foreach
// find
// find_if
