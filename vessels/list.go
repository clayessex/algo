package vessels

type ListNode[T any] struct {
	next  *ListNode[T]
	prev  *ListNode[T]
	value T
}

func (n *ListNode[T]) Next() *ListNode[T] {
	if n == nil {
		return nil
	}
	return n.next
}

func (n *ListNode[T]) Prev() *ListNode[T] {
	if n == nil {
		return nil
	}
	return n.prev
}

/** insert n before p */
func (n *ListNode[T]) insertBefore(p *ListNode[T]) *ListNode[T] {
	if n == nil || p == nil {
		panic("ListNode insertBefore invalid nil node")
	}
	if p.prev != nil {
		p.prev.next = n
	}
	n.prev = p.prev
	n.next = p
	p.prev = n
	return n
}

/** insert n after p */
func (n *ListNode[T]) insertAfter(p *ListNode[T]) *ListNode[T] {
	if n == nil || p == nil {
		panic("ListNode insertAfter invalid nil node")
	}
	if p.next != nil {
		p.next.prev = n
	}
	n.next = p.next
	n.prev = p
	p.next = n
	return n
}

func (n *ListNode[T]) remove() *ListNode[T] {
	if n == nil {
		panic("ListNode remove invalid nil node")
	}
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
	n.prev = nil // prevent memory leak (maybe)
	n.next = nil
	return n
}

type List[T any] struct {
	head *ListNode[T] // first node
	tail *ListNode[T] // last node
	len  int
}

func NewList[T any]() *List[T] {
	return &List[T]{nil, nil, 0}
}

func (l *List[T]) Len() int {
	return l.len
}

// insert before pos or append if pos is nil, return new node
func (l *List[T]) insert(v T, pos *ListNode[T]) *ListNode[T] {
	n := &ListNode[T]{nil, nil, v}

	if pos == nil { // append
		if l.tail != nil {
			n.insertAfter(l.tail)
		}
	} else {
		n.insertBefore(pos)
	}
	if n.Prev() == nil {
		l.head = n
	}
	if n.Next() == nil {
		l.tail = n
	}
	l.len++
	return n
}

// remove pos from list, returns removed node
func (l *List[T]) remove(pos *ListNode[T]) *ListNode[T] {
	if pos == nil {
		panic("List remove invalid pos")
	}
	if l.head == pos {
		l.head = pos.Next()
	}
	if l.tail == pos {
		l.tail = pos.Prev()
	}
	pos.remove()
	l.len--
	return pos
}

func (l *List[T]) InsertAfter(v T, pos *ListNode[T]) *ListNode[T] {
	return l.insert(v, pos.Next())
}

func (l *List[T]) InsertBefore(v T, pos *ListNode[T]) *ListNode[T] {
	return l.insert(v, pos)
}

func (l *List[T]) PushBack(v T) *ListNode[T] {
	return l.insert(v, nil)
}

func (l *List[T]) PushFront(v T) *ListNode[T] {
	return l.insert(v, l.head)
}

func (l *List[T]) PopBack() T {
	return l.remove(l.tail).value
}

func (l *List[T]) PopFront() T {
	return l.remove(l.head).value
}

func (l *List[T]) Append(v T) *ListNode[T] {
	return l.PushBack(v)
}

func (l *List[T]) Front() *ListNode[T] {
	return l.head
}

func (l *List[T]) Back() *ListNode[T] {
	return l.tail
}

func (l *List[T]) Clear() {
	for l.Len() > 0 {
		l.remove(l.tail)
	}
}

func (l *List[T]) At(index int) T {
	if index < 0 || index >= l.Len() {
		panic("List At() index out of range")
	}

	p := l.head
	for i := 0; i < index; i++ {
		p = p.Next()
	}
	return p.value
}

// TODO: iterations
// foreach
// find
// find_if
