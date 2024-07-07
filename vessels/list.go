package vessels

type ListNode[T any] struct {
	next  *ListNode[T]
	prev  *ListNode[T]
	value T
}

type List[T any] struct {
	head *ListNode[T] // first node
	tail *ListNode[T] // last node
	len  int
}

func NewList[T any]() *List[T] {
	return &List[T]{nil, nil, 0}
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

func (l *List[T]) Len() int {
	return l.len
}

// insert before pos, return new node
func (l *List[T]) insert(v T, pos *ListNode[T]) *ListNode[T] {
	right := pos

	var left *ListNode[T] = nil
	if pos == nil {
		left = l.tail
	} else {
		left = pos.Prev()
	}

	n := &ListNode[T]{right, left, v}

	if left == nil {
		l.head = n
	} else {
		left.next = n
	}

	if right == nil {
		l.tail = n
	} else {
		right.prev = n
	}

	l.len++
	return n
}

// returns removed node
func (l *List[T]) remove(pos *ListNode[T]) *ListNode[T] {
	if pos == nil {
		panic("invalid pos deleting from List")
	}

	left := pos.Prev()
	right := pos.Next()

	// prevent memory leaks (maybe)
	pos.next = nil
	pos.prev = nil

	if left == nil {
		l.head = right
	} else {
		left.next = right
	}

	if right == nil {
		l.tail = left
	} else {
		right.prev = left
	}

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
