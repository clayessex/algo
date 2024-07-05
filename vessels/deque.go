package vessels

/** Queue */
const INITIAL_DEQUE_SIZE = 32

type Deque[T any] struct {
	buf []T

	/** capacity is cap - 1 */
	cap int

	/** head is the next avilable empty slot :: back/end */
	head int

	/** tail is the last value unless tail == head :: front/begin */
	tail int
}

func NewDeque[T any](size ...int) *Deque[T] {
	sz := INITIAL_DEQUE_SIZE
	if len(size) >= 1 {
		sz = size[0]
	}
	return &Deque[T]{
		make([]T, sz),
		sz,
		0,
		0,
	}
}

func (d *Deque[T]) copy(dst []T) int {
	sz := min(d.Len(), len(dst))
	p := d.tail
	for i := 0; i < sz; i++ {
		dst[i] = d.buf[p]
		p = d.next(p)
	}
	return sz
}

/** double the size of the buffer and copy the old one over */
func (d *Deque[T]) grow() {
	newSize := d.cap * 2
	newBuf := make([]T, newSize)
	d.head = d.copy(newBuf)
	d.tail = 0
	d.buf = newBuf
	d.cap = newSize
}

/** half the size of the buffer if the current Len will fit */
func (d *Deque[T]) shrink() {
	newSize := len(d.buf) / 2
	if newSize < INITIAL_DEQUE_SIZE || newSize <= d.Len() {
		return
	}
	newBuf := make([]T, newSize)
	d.head = d.copy(newBuf)
	d.tail = 0
	d.buf = newBuf
	d.cap = newSize
}

/** calculate the next index in sequence, does not detect a full buffer */
func (d *Deque[T]) next(index int) int {
	return (index + 1) % d.cap
}

/** calculate the previous index in sequence, does not detect an empty buffer */
func (d *Deque[T]) prev(index int) int {
	if index == 0 {
		return d.cap - 1
	}
	return index - 1
}

func (d *Deque[T]) Len() int {
	if d.tail <= d.head {
		return d.head - d.tail
	}
	return d.cap - (d.tail - d.head)
}

func (d *Deque[T]) Cap() int {
	return d.cap
}

func (d *Deque[T]) PushBack(v T) {
	if d.Len() == d.cap-1 {
		d.grow()
	}
	d.buf[d.head] = v
	d.head = d.next(d.head)
}

func (d *Deque[T]) PushFront(v T) {
	if d.Len() == d.cap-1 {
		d.grow()
	}
	d.tail = d.prev(d.tail)
	d.buf[d.tail] = v
}

func (d *Deque[T]) PopBack() T {
	if d.Len() == 0 {
		panic("can't PopBack() from an empty deque")
	}
	d.head = d.prev(d.head)
	return d.buf[d.head]
}

func (d *Deque[T]) PopFront() T {
	if d.Len() == 0 {
		panic("can't PopFront() from an empty deque")
	}
	result := d.buf[d.tail]
	d.tail = d.next(d.tail)
	return result
}

func (d *Deque[T]) Clear() {
	d.tail = 0
	d.head = 0
}

func (d *Deque[T]) Shrink() {
	d.shrink()
}

func (d *Deque[T]) Clone() *Deque[T] {
	clone := &Deque[T]{
		make([]T, d.cap),
		d.cap,
		0,
		0,
	}
	clone.head = d.copy(clone.buf)
	return clone
}
