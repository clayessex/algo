package vessels

// OrderedMap is a map that remembers the insertion order of elements. All operations
// are O(1) except the At() function, which is O(N).
type OrderedMap[K comparable, V any] struct {
	data  map[K]V            // map keys to values
	nodes map[K]*ListNode[K] // map keys to ListNodes
	ord   List[K]            // list order of keys
}

// Create a new OrderedMap with an optional initial allocation size
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

// Returns the current length of the map
func (m *OrderedMap[K, V]) Len() int {
	return len(m.data)
}

// Returns true if the OrderedMap contains the given key
func (m *OrderedMap[K, V]) Contains(key K) bool {
	_, ok := m.data[key]
	return ok
}

// Returns the value for the given key if the key exists, otherwise a default
// initialized value and false
func (m *OrderedMap[K, V]) Value(key K) (V, bool) {
	v, ok := m.data[key]
	return v, ok
}

// Insert the key/value pair into the map (same as Push())
func (m *OrderedMap[K, V]) Insert(key K, value V) {
	if !m.Contains(key) { // if exists, overwrite in place
		m.ord.PushBack(key)
		m.nodes[key] = m.ord.End().Prev()
	}
	m.data[key] = value
}

// Delete the key/value pair for the given key
func (m *OrderedMap[K, V]) Delete(key K) bool {
	n, ok := m.nodes[key]
	if !ok {
		return false
	}
	m.ord.RemoveNode(n)
	delete(m.nodes, key)
	delete(m.data, key)
	return true
}

// Push the key/value pair into the map (same as Insert())
func (m *OrderedMap[K, V]) Push(key K, value V) {
	m.Insert(key, value)
}

// Pop the last key/value pair from the end of the map unless the map is empty
// and then it returns a default initialized value and false
func (m *OrderedMap[K, V]) Pop() (K, bool) {
	key, ok := m.ord.PopBack()
	if !ok { // only happens when list is empty
		return *new(K), false
	}
	delete(m.nodes, key)
	delete(m.data, key)
	return key, true
}

// Returns the key following the given key
func (m *OrderedMap[K, V]) Next(key K) (K, bool) {
	n, ok := m.nodes[key]
	if !ok {
		return *new(K), false
	}
	nextNode := n.Next()
	if nextNode == m.ord.End() {
		return *new(K), false
	}

	return nextNode.value, true
}

// Returns the key preceding the given key if there is one, otherwise returns a
// default initialized value and false
func (m *OrderedMap[K, V]) Prev(key K) (K, bool) {
	n, ok := m.nodes[key]
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

// Returns the oldest key in insertion order unless the map is empty, then it
// returns a default initialized value and false
func (m *OrderedMap[K, V]) First() (K, bool) {
	if m.ord.Len() == 0 {
		return *new(K), false
	}
	return m.ord.Front()
}

// Returns the newest key in insertion order unless the map is empty, then it
// returns a default initialized value and false
func (m *OrderedMap[K, V]) Last() (K, bool) {
	if m.ord.Len() == 0 {
		return *new(K), false
	}
	return m.ord.Back()
}

// Remove all key/value pairs from the map
func (m *OrderedMap[K, V]) Clear() {
	clear(m.data)
	clear(m.nodes)
	m.ord.Clear()
}

// Return a slice containing all of the keys in insertion order
func (m *OrderedMap[K, V]) Keys() []K {
	return m.ord.Values()
}

// Return a slice containing all of the values in insertion order
func (m *OrderedMap[K, V]) Values() []V {
	r := make([]V, 0, len(m.data))
	m.ord.Range(func(key K) {
		r = append(r, m.data[key])
	})
	return r
}

// Return the value at position index in insertion order. If the index is out
// of bounds then return a default initialized value and false. The index
// lookup is an O(N) operation.
func (m *OrderedMap[K, V]) At(index int) (V, bool) {
	key, ok := m.ord.At(index)
	if !ok {
		return *new(V), false
	}
	value, ok := m.data[key]
	return value, ok
}

// Iterate over the map in insertion order calling function f on each key/value
// pair
func (m *OrderedMap[K, V]) Range(f func(K, V)) {
	m.ord.Range(func(key K) {
		value := m.data[key]
		f(key, value)
	})
}
