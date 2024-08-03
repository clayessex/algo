package vessels

type Stack[T any] Deque[T]

func NewStack[T any](size ...int) *Stack[T] {
	return (*Stack[T])(NewDeque[T](size...))
}

func (s *Stack[T]) Cap() int {
	return (*Deque[T])(s).Cap()
}

func (s *Stack[T]) Len() int {
	return (*Deque[T])(s).Len()
}

func (s *Stack[T]) Push(v T) {
	(*Deque[T])(s).PushBack(v)
}

func (s *Stack[T]) Pop() (T, bool) {
	return (*Deque[T])(s).PopBack()
}

func (s *Stack[T]) At(index int) (T, bool) {
	return (*Deque[T])(s).At(index)
}

func (s *Stack[T]) Clear() {
	(*Deque[T])(s).Clear()
}

func (s *Stack[T]) Clone() *Stack[T] {
	return (*Stack[T])((*Deque[T])(s).Clone())
}
