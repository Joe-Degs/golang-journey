package ringbuffer

import (
	"testing"
)

func TestType(t *testing.T) {
	buf := New(2)
	buf.Add(1, 2)
	buf.Peek()
	if buf.idx != -1 {
		t.Error("buffer idx doesnt work")
	}

	buf.Next()
	if buf.idx != 0 {
		t.Error("buffer read is flawed.")
	}
}

func TestBufferWrite(t *testing.T) {
	buf := New(2)
	err := buf.Add(1, 3, 4, 5, 6)
	if err == nil {
		t.Error("buffer should return error")
	}
	t.Log(err)

	if n := buf.insert_idx(); n != 0 {
		t.Error("something wrong somewhere.")
	}
}

func TestBufferRead(t *testing.T) {
	buf := New(4)
	if n := buf.read_idx(); n != 0 {
		t.Error("fuck!")
	}
}
