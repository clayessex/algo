/**
* Deque
* Double ended queue implemented using a ring buffer
* Holds (size - 1) items of T before the size is automatically doubled.
 */

package vessels

const INITIAL_DEQUE_SIZE = 32

/** Deque */
type Deque[T any] struct {
	/**
	 * buf holds the deque items of T - len(buf) is also the
	 * capacity of the Deque
	 */
	buf []T

	/**
	 * head is the next avilable empty slot :: back/end
	 * PushBack/PopBack insert here
	 */
	head int

	/**
	 * tail is the last element unless tail == head :: front/begin
	 * PushFront/PopFront insert before here
	 */
	tail int
}

/**
 * Create a new Deque[T] sized to hold size-1 items
 * before automatically resizing.
 */
func NewDeque[T any](size ...int) *Deque[T] {
	sz := INITIAL_DEQUE_SIZE
	if len(size) >= 1 {
		sz = size[0]
	}
	return &Deque[T]{
		make([]T, sz),
		0,
		0,
	}
}

/**
 * copy the deque buffer into dst, panics if dst is not large enough
 * returns number of elements copied
 */
func (d *Deque[T]) copy(dst []T) int {
	if len(dst) < d.Len() {
		panic("deque too large to copy to dst")
	}
	var n int
	if d.tail <= d.head {
		n = copy(dst, d.buf[d.tail:d.head])
	} else {
		n = copy(dst, d.buf[d.tail:])
		n += copy(dst[n:], d.buf[:d.head])
	}
	return n
}

/** double the size of the buffer and copy the old one over */
func (d *Deque[T]) grow() {
	newSize := len(d.buf) * 2
	newBuf := make([]T, newSize)
	d.head = d.copy(newBuf)
	d.tail = 0
	d.buf = newBuf
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
}

/** calculate the next index in sequence, does not detect a full buffer */
func (d *Deque[T]) next(index int) int {
	return (index + 1) % len(d.buf)
}

/** calculate the previous index in sequence, does not detect an empty buffer */
func (d *Deque[T]) prev(index int) int {
	if index == 0 {
		return len(d.buf) - 1
	}
	return index - 1
}

/** length of the deque in use */
func (d *Deque[T]) Len() int {
	if d.tail <= d.head {
		return d.head - d.tail
	}
	return len(d.buf) - (d.tail - d.head)
}

/** capacity of the deque */
func (d *Deque[T]) Cap() int {
	return len(d.buf)
}

/** append to the end of the buffer */
func (d *Deque[T]) PushBack(v T) {
	if d.Len() == len(d.buf)-1 {
		d.grow()
	}
	d.buf[d.head] = v
	d.head = d.next(d.head)
}

/** insert before the beginning of the buffer */
func (d *Deque[T]) PushFront(v T) {
	if d.Len() == len(d.buf)-1 {
		d.grow()
	}
	d.tail = d.prev(d.tail)
	d.buf[d.tail] = v
}

/** remove and return the last element */
func (d *Deque[T]) PopBack() T {
	if d.Len() == 0 {
		panic("can't PopBack() from an empty deque")
	}
	d.head = d.prev(d.head)
	return d.buf[d.head]
}

/** remove and return the first element */
func (d *Deque[T]) PopFront() T {
	if d.Len() == 0 {
		panic("can't PopFront() from an empty deque")
	}
	result := d.buf[d.tail]
	d.tail = d.next(d.tail)
	return result
}

/**
 * return the first element from the beginning of the Deque without removing it,
 * panics if the Deque is empty
 */
func (d *Deque[T]) Front() T {
	if d.Len() == 0 {
		panic("can't get front() element from an empty deque")
	}
	return d.buf[d.tail]
}

/**
 * return the last element from the end of the Deque without removing it,
 * panics if the Deque is empty
 */
func (d *Deque[T]) Back() T {
	if d.Len() == 0 {
		panic("can't get back() element from an empty deque")
	}
	return d.buf[d.prev(d.head)]
}

/**
 * return the element at index without removing it,
 * panics if index is outside of range [0:Len()]
 */
func (d *Deque[T]) At(index int) T {
	if index < 0 || index >= d.Len() {
		panic("Deque At() invalid index")
	}
	offset := (d.tail + index) % len(d.buf)
	return d.buf[offset]
}

/** remove all elements, leaving the deque empty */
func (d *Deque[T]) Clear() {
	d.tail = 0
	d.head = 0
}

/** reduce the capacity by half unless the current elements won't fit */
func (d *Deque[T]) Shrink() {
	d.shrink()
}

/** create a clone of the deque */
func (d *Deque[T]) Clone() *Deque[T] {
	clone := &Deque[T]{
		make([]T, len(d.buf)),
		0,
		0,
	}
	clone.head = d.copy(clone.buf)
	return clone
}
