package ringbuffer

import (
	"errors"
)

// this implementation looks horrible, mayber i'll plan
// and do it again sometime.

var (
	ErrBufull error = errors.New("buffer is full, consider reading before writing")
	//typeMismatch string = `expected %s but got %s`
)

type RingBuffer struct {
	N int // size of buffer

	counter int  // keeps track of space left in buffer.
	idx     int  // the last read index.
	flag    bool // true -> back of insert idx , false -> front of insert idx

	// buf stores the actual values of the buffer
	// it only does so after making sure the value and t are of the same type.
	buf []interface{}
}

func (r *RingBuffer) init() {
	r.buf = make([]interface{}, r.N)
}

func (r *RingBuffer) Add(items ...interface{}) error {
	if r.counter < len(items) {
		return errors.New("can't fit items into buffer.")
	}
	// prevent overriding unread elements.

	for i := 0; i < len(items); i++ {
		r.buf[r.insert_idx()] = items[i]
	}
	return nil
}

// next_idx returns next index to write to.
func (r *RingBuffer) insert_idx() int {
	// counter is the amt of space left in buffer.
	var idx int
	if r.counter == 0 {
		// reset counter
		r.counter = r.N
	}
	r.counter -= 1
	idx = r.N - 1 - r.counter
	return idx
}

// returns the next index to read to.
func (r *RingBuffer) read_idx() int {
	if r.idx == r.N-1 {
		// reset idx to start of buffer
		r.idx = -1
	}

	// return previously read buffer position
	// if next position insert_position.
	if r.idx+1 == r.insert_idx() {
		// reset idx's
		r.counter += 1
		return r.idx
	}
	r.idx += 1
	return r.idx
}

// IncreaseCap increases the capacity of the buffer.
func (r *RingBuffer) IncreaseCap() {
	return
}

// Peek peeps the next item in the buffer.
func (r *RingBuffer) Peek() interface{} {
	//read, send idx back.
	p := r.buf[r.read_idx()]
	r.idx -= 1
	return p

}

// Next reads the next item in the buffer.
func (r *RingBuffer) Next() interface{} { return r.buf[r.read_idx()] }

// Read peeps at an arbitrary buffer index, returning it.
func (r *RingBuffer) Read(index int) interface{} {
	return 0
}

func New(size int) *RingBuffer {
	buf := &RingBuffer{
		N:       size,
		counter: size,
		idx:     0,
	}
	buf.init()
	return buf
}
