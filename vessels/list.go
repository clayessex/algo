package vessels

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

/** Moves [first, last) before pos */
func splice[T any](pos, first, last *ListNode[T]) {
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
	o.remove().insertBefore(n)
	if tmp != n {
		n.remove().insertBefore(tmp)
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

func (list *List[T]) RemoveNode(pos *ListNode[T]) *ListNode[T] {
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
	return list.RemoveNode(list.head.prev).value
}

func (list *List[T]) PopFront() T {
	if list.len == 0 {
		panic("list PopFront() called on empty list")
	}
	return list.RemoveNode(list.head.next).value
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
	p := list.End()
	for i := 0; i < list.len; i++ {
		p = list.Begin().remove().insertBefore(p)
	}
}

func (list *List[T]) Splice(pos, first, last *ListNode[T]) {
	splice(pos, first, last)
}

func ListRemoveFunc[T any](list *List[T], v T, comp func(a, b T) bool) {
	if list.Len() == 0 {
		return
	}

	p := list.Begin()
	for p != list.End() {
		if comp(p.value, v) {
			tmp := p
			p = p.Next()
			tmp.remove()
		} else {
			p = p.Next()
		}
	}
}

func ListRemove[T comparable](list *List[T], v T) {
	ListRemoveFunc(list, v, func(a, b T) bool {
		return a == b
	})
}

// func Remove[T any](list *List[T}, v T) {
//
// }

// TODO: iterations
// foreach
// find
// find_if
