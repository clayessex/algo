package vessels

type OrderedMap[K comparable, V any] struct {
	data map[K]V
	del  map[K]*ListNode[K]
	ord  List[K]
}

func NewOrderedMap[K comparable, V any](size ...int) *OrderedMap[K, V] {
	sz := INITIAL_DEQUE_SIZE
	if len(size) > 0 {
		sz = size[0]
	}
	r := OrderedMap[K, V]{
		make(map[K]V, sz),
		make(map[K]*ListNode[K], sz),
		*NewList[K](),
	}
	return &r
}

func (m *OrderedMap[K, V]) Len() int {
	return len(m.data)
}

func (m *OrderedMap[K, V]) Contains(key K) bool {
	_, ok := m.data[key]
	return ok
}

func (m *OrderedMap[K, V]) Value(key K) (V, bool) {
	v, ok := m.data[key]
	return v, ok
}

func (m *OrderedMap[K, V]) Insert(key K, value V) {
	if !m.Contains(key) { // if exists, overwrite in place
		m.ord.PushBack(key)
		m.del[key] = m.ord.End()
	}
	m.data[key] = value
}

func (m *OrderedMap[K, V]) Delete(key K) bool {
	n, ok := m.del[key]
	if !ok {
		return false
	}
	m.ord.RemoveNode(n)
	delete(m.del, key)
	delete(m.data, key)
	return true
}

func (m *OrderedMap[K, V]) Push(key K, value V) {
	m.Insert(key, value)
}

func (m *OrderedMap[K, V]) Pop() K {
	if m.ord.Len() == 0 {
		panic("OrderedMap can't Pop() when empty")
	}
	key := m.ord.PopBack()
	delete(m.del, key)
	delete(m.data, key)
	return key
}

func (m *OrderedMap[K, V]) Next(key K) (K, bool) {
	n, ok := m.del[key]
	if !ok {
		var zero K
		return zero, false
	}
	nextNode := n.Next()
	if nextNode == m.ord.End() {
		var zero K
		return zero, false
	}

	return nextNode.value, true
}

func (m *OrderedMap[K, V]) Prev(key K) (K, bool) {
	n, ok := m.del[key]
	if !ok {
		var zero K
		return zero, false
	}
	if n == m.ord.Begin() {
		var zero K
		return zero, false
	}
	return n.Prev().value, true
}

func (m *OrderedMap[K, V]) First() K {
	if m.ord.Len() == 0 {
		panic("OrderedMap can't return First() when empty")
	}
	return m.ord.Back()
}

func (m *OrderedMap[K, V]) Last() K {
	if m.ord.Len() == 0 {
		panic("OrderedMap can't return Last() when empty")
	}
	return m.ord.Front()
}

func (m *OrderedMap[K, V]) Clear() {
	clear(m.data)
	clear(m.del)
	m.ord.Clear()
}

func (m *OrderedMap[K, V]) Keys() []K {
	return m.ord.Values()
}

func (m *OrderedMap[K, V]) Values() []V {
	r := make([]V, 0, len(m.data))
	m.ord.Range(func(key K) {
		r = append(r, m.data[key])
	})
	return r
}

func (m *OrderedMap[K, V]) Range(f func(K, V)) {
	m.ord.Range(func(key K) {
		value := m.data[key]
		f(key, value)
	})
}
