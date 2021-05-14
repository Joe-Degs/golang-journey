package contextimpl

import (
	"errors"
	"sync"
	"time"
)

// its nice that context itself is an interface..
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

type emptyCtx int

// can't use empty structs for this type of stuff because tests are not passing.
// the empty struct have zero size at runtime so, they can basically have the same
// pointer location is the compiler does it optimizations and stuff

func (emptyCtx) Deadline() (time.Time, bool)       { return time.Time{}, false }
func (emptyCtx) Done() <-chan struct{}             { return nil }
func (emptyCtx) Err() error                        { return nil }
func (emptyCtx) Value(key interface{}) interface{} { return nil }

var (
	// this is to make sure when the context package is called multiple time
	// it does not allocate too many background and todo contexts for use.
	// no matter how many times you call this package, it'll always return
	// this two contexts for use.
	// at the end of the day you will only allocate two ints ever if you use
	// the context package.
	todo = new(emptyCtx)
	bg   = new(emptyCtx)
)

func Background() Context { return bg }
func TODO() Context       { return todo }

// implementing WithCancel() function.
// it returns a Context and CancelFunc
// a CancelFunc allows to cancel the context after WithCancel has been called.

type cancelCtx struct {
	Context
	done chan struct{}
	err  error
	mu   sync.Mutex
}

func (ctx *cancelCtx) Done() <-chan struct{} {
	return ctx.done
}

func (ctx *cancelCtx) Err() error {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	return ctx.err
}

var Canceled = errors.New("context canceled") // returned by Context.Err
type CancelFunc func()

func WithCancel(parent Context) (Context, CancelFunc) {
	ctx := &cancelCtx{
		Context: parent,
		done:    make(chan struct{}),
	}

	cancel := func() { ctx.cancel(Canceled) }

	go func() {
		select {
		case <-parent.Done():
			ctx.cancel(parent.Err())
		case <-ctx.Done():
		}
	}()

	return ctx, cancel
}

// this takes an error and just cancels the context with the error.
func (ctx *cancelCtx) cancel(err error) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	if ctx.err != nil {
		return
	}
	ctx.err = err
	close(ctx.done)
}

// WithDeadline
// cancels context after given time
// and sets ctx.err to DeadlineExceeded
var DeadlineExceeded = errors.New("deadline exceeded")

type deadlineCtx struct {
	*cancelCtx
	deadline time.Time
}

func (ctx *deadlineCtx) Deadline() (time.Time, bool) {
	return ctx.deadline, true
}

func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc) {
	cctx, cancel := WithCancel(parent)

	ctx := &deadlineCtx{
		cancelCtx: cctx.(*cancelCtx),
		deadline:  deadline,
	}

	t := time.AfterFunc(time.Until(deadline), func() {
		ctx.cancel(DeadlineExceeded)
	})

	//timer returns a timer, so we stop that timer and cancel if CancelFunc is
	// called.
	stop := func() {
		t.Stop()
		cancel()
	}

	return ctx, stop
}

//WithTimeout
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
}
