package main

/** Queue */
const INITIAL_DEQUE_SIZE = 32

type Deque struct {
	buf []int

	/** capacity is cap - 1 */
	cap int

	/** head is the next avilable empty slot :: back/end */
	head int

	/** tail is the last value unless tail == head :: front/begin */
	tail int
}

func MakeDeque() *Deque {
	return MakeDequeSized(INITIAL_DEQUE_SIZE)
}

func MakeDequeSized(size int) *Deque {
	return &Deque{
		make([]int, size),
		size,
		0,
		0,
	}
}

func (d *Deque) grow() {
	newSize := d.cap * 2
	newBuf := make([]int, newSize)
	newHead := 0
	for i := d.tail; i != d.head; i = (i + 1) % d.cap {
		newBuf[newHead] = d.buf[i]
		newHead++
	}
	d.buf = newBuf
	d.cap = newSize
	d.tail = 0
	d.head = newHead
}

func (d *Deque) Size() int {
	// 0 1 2 3 4 : N==5
	// T x x x H == 4 - 0 = 4 : always (cap - 1)
	// . T x H . == 3 - 1 = 2 :
	if d.tail <= d.head {
		return d.head - d.tail
	}

	// 0 1 2 3 4 : N==5
	// x H . T x == 5 - (3 - 1) = 3
	// H . . . T == 5 - (4 - 0) = 1
	return d.cap - (d.tail - d.head)
}

func (d *Deque) Cap() int {
	return d.cap
}

func (d *Deque) PushBack(v int) {
	if d.Size() == d.cap-1 {
		d.grow()
	}
	d.buf[d.head] = v
	d.head = (d.head + 1) % d.cap
}

func (d *Deque) PushFront(v int) {
	if d.Size() == d.cap-1 {
		d.grow()
	}
	d.tail = (d.tail - 1) % d.cap
	d.buf[d.tail] = v
}

func (d *Deque) PopBack() int {
	if d.Size() == 0 {
		panic("can't PopBack() from an empty deque")
	}
	d.head = (d.head - 1) % d.cap
	return d.buf[d.head]
}

func (d *Deque) PopFront() int {
	if d.Size() == 0 {
		panic("can't PopFront() from an empty deque")
	}
	result := d.buf[d.tail]
	d.tail = (d.tail + 1) % d.cap
	return result
}

func (d *Deque) Clear() {
	d.tail = 0
	d.head = 0
}

func (d *Deque) Clone() *Deque {
	clone := &Deque{
		make([]int, d.cap),
		d.cap,
		d.head,
		d.tail,
	}
	copy(clone.buf, d.buf)
	return clone
}
