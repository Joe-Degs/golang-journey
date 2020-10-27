// Package runner manages the running and lifetime of a process
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner runs a set of tasks within a given timeout and can be
// shutdown on a operating system interrupt
type Runner struct {
	// interrupt channel reports a signal from the operating system
	interrupt chan os.Signal

	// complete channeln reports that processing is done
	complete chan error

	// timeout reports that time has run out
	timeout <-chan time.Time

	// tasks holds a set of functions that are executed
	// syncronously in index order
	tasks []func(int)
}

// ErrTimeout is returned when a value is recieved on the timeout
var ErrTimeout = errors.New("recieved timeout")

// ErrInterrupt is returned when an event from the OS is recieved
var ErrInterrupt = errors.New("recieved interrupt")

// New returns a ready to use runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add attaches tasks to the runner.
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start runs all tasks and monitors channel events
func (r *Runner) Start() error {
	// We want to recieve all interrupt based signals
	signal.Notify(r.interrupt, os.Interrupt)

	// run the different tasks on a different goroutine
	go func() {
		r.complete <- r.run()
	}()

	select {
	// signalled is processing is done.
	case err := <-r.complete:
		return err

	// signalled when when we run out of time
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		// check for an interrupt signal from the OS
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// execute the registered task
		task(id)
	}

	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	// signaled when an interrupt event is sent.
	case <-r.interrupt:
		// stop recieving any further signals
		signal.Stop(r.interrupt)
		return true

	// continue running as normal.
	default:
		return false
	}
}
