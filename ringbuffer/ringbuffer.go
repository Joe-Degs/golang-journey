package ringbuffer

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrBufull    error  = errors.New("buffer is full, consider reading before writing")
	typeMismatch string = `expected %s but got %s`
)

type RingBuffer struct {
	size int // reps size of buffer
	//t    string      // reps the static type of buffer
	static_type string // hold sample type for checking validity of other type.

	// buf stores the actual values of the buffer
	// it only does so after making sure the value and t are of the same type.
	buf []interface{}
}

// get the static type of the parameter passed in
// as a string. mostly in the form <packagename.Type>
func getType(t interface{}) string {
	s := fmt.Sprintf("%T", t)
	if !strings.Contains(s, "main") {
		return s
	}
	return s[strings.LastIndex(s, ".")+1:]
}

// initialize the slice to the desired size
func (r *RingBuffer) init() {
	r.buf = make([]interface{}, r.size)
}

func (r *RingBuffer) Add(element ...interface{}) {
	for _, val := range element {
		if t := getType(val); t != r.static_type {
			panic(fmt.Errorf(typeMismatch, r.static_type, t))
		}
	}
}

func (r *RingBuffer) ReadNext() interface{} {
	return 0
}

func (r *RingBuffer) Read(index int) interface{} {
	return 0
}

// st -> static type of the underlying buffer.
// size is the size of underlying buffer.
// for now i think the best way is to pass a sample type
// of type of values to store in the underlying buffer.
func New(size int, st interface{}) *RingBuffer {
	buf := &RingBuffer{
		size:        size,
		static_type: getType(st),
	}
	buf.init()
	return buf
}
