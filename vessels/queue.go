package vessels

type Queue[T any] Deque[T]

func NewQueue[T any](size ...int) *Queue[T] {
	return (*Queue[T])(NewDeque[T](size...))
}

func (q *Queue[T]) Len() int {
	return (*Deque[T])(q).Len()
}

func (q *Queue[T]) Cap() int {
	return (*Deque[T])(q).Cap()
}

func (q *Queue[T]) Push(v T) {
	(*Deque[T])(q).PushBack(v)
}

func (q *Queue[T]) Pop() T {
	return (*Deque[T])(q).PopFront()
}

func (q *Queue[T]) Clear() {
	(*Deque[T])(q).Clear()
}

func (q *Queue[T]) Clone() *Queue[T] {
	return (*Queue[T])((*Deque[T])(q).Clone())
}
