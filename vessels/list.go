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

func (n *ListNode[T]) remove() *ListNode[T] {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.prev = n
	n.next = n
	return n
}

func (n *ListNode[T]) Swap(o *ListNode[T]) {
	n.value, o.value = o.value, n.value
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

func (list *List[T]) remove(pos *ListNode[T]) *ListNode[T] {
	list.len--
	return pos.remove()
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
	return list.remove(list.head.prev).value
}

func (list *List[T]) PopFront() T {
	if list.len == 0 {
		panic("list PopFront() called on empty list")
	}
	return list.remove(list.head.next).value
}

func (list *List[T]) Append(v T) *ListNode[T] {
	return list.PushBack(v)
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
	p := list.End()
	for i := 0; i < list.len; i++ {
		p = list.Begin().remove().insertBefore(p)
	}
}

func splice[T any](pos *ListNode[T], first *ListNode[T], last *ListNode[T]) *ListNode[T] {
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

	return oright
}

/** merge two sorted lists, [first2, last2) into [first1, last1) */
func merge[T cmp.Ordered](
	first1 *ListNode[T], last1 *ListNode[T],
	first2 *ListNode[T], last2 *ListNode[T],
) *ListNode[T] {
	newBegin := first1 // TODO: subopt
	if first2.value < first1.value {
		newBegin = first2
	}

	for first1 != last1 && first2 != last2 {
		if first2.value < first1.value {
			next := first2.next
			splice(first1, first2, next)
			first2 = next
		} else {
			first1 = first1.next
		}
	}

	if first2 != last2 {
		splice(last1, first2, last2)
	}

	return newBegin
}

// TODO: iterations
// foreach
// find
// find_if
