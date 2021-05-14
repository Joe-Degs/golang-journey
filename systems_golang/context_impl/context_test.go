package contextimpl

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestBackgroundNotTODO(t *testing.T) {
	todo := fmt.Sprintf("%q", TODO())
	bg := fmt.Sprintf("%q", Background())

	if todo == bg {
		t.Errorf("TODO and Background are equal: %q vs %q", todo, bg)
	}
}

func TestWithCancel(t *testing.T) {
	ctx, cancel := WithCancel(Background())

	if err := ctx.Err(); err != nil {
		t.Errorf("error should be nil first, got %q", err)
	}
	cancel()

	<-ctx.Done()
	if err := ctx.Err(); err != Canceled {
		t.Errorf("error should be cancelled now. got %v", err)
	}
}

func TestWithCancelConcurrent(t *testing.T) {
	ctx, cancel := WithCancel(Background())

	time.AfterFunc(time.Second*1, cancel)

	if err := ctx.Err(); err != nil {
		t.Errorf("error should be nil first, got %v", err)
	}

	<-ctx.Done()
	if err := ctx.Err(); err != Canceled {
		t.Errorf("error should be canceled now, got %v", err)
	}
}

func TestWithCancelPropagation(t *testing.T) {
	ctxA, cancelA := WithCancel(Background())
	ctxB, cancelB := WithCancel(ctxA)
	defer cancelB()

	cancelA()

	select {
	case <-ctxB.Done():
	case <-time.After(1 * time.Second):
		t.Errorf("time out")
	}

	if err := ctxB.Err(); err != Canceled {
		t.Errorf("error should be canceled now, got %v", err)
	}
}

func TestWithDeadline(t *testing.T) {
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := WithDeadline(Background(), deadline)

	if d, ok := ctx.Deadline(); !ok || d != deadline {
		t.Errorf("expected deadline %v, got %v", deadline, d)
	}

	then := time.Now()
	<-ctx.Done()
	if d := time.Since(then); math.Abs(d.Seconds()-2.0) > 0.1 {
		t.Errorf("should have been done after 2.0 seconds, took %v", d)
	}
	if err := ctx.Err(); err != DeadlineExceeded {
		t.Errorf("expected deadline exceeded, got %v", err)
	}

	cancel()
	if err := ctx.Err(); err != DeadlineExceeded {
		t.Errorf("expected deadline exceeded, got %v", err)
	}
}
