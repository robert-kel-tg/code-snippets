package concurency

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	// interrupt channel reports a signal from os
	interrupt chan os.Signal

	// complete channel reports that processing is done
	complete chan error

	// timeout reports that has run out
	timeout <-chan time.Time

	// holds a set of functions
	tasks []func(int)
}

// ErrTimeout when a value is received on timeout
var ErrTimeout = errors.New("Received timeout")

// ErrInterrupt when an event form OS is received
var ErrInterrupt = errors.New("Received interrupt")

// https://www.safaribooksonline.com/library/view/go-in-action/9781617291784/kindle_split_015.html
// Returns new ready-to-use Runner
func NewRunner(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Attached tasks to the runner. Task is an function which takes an int ID
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {

	// Subscribe to receive all interrupt signals from OS
	signal.Notify(r.interrupt, os.Interrupt)

	// Run every new task on single but separate from main goroutine
	go func(ir *Runner) {
		ir.complete <- r.run()
	}(r)

	select {
	// Signalled when processing is done
	case err := <-r.complete:
		return err
	// Signalled when we run out of time
	case <-r.timeout:
		return ErrTimeout
	}
}

// Executes each registered task
func (r *Runner) run() error {

	// Check for interrupt signal from the OS
	if r.gotInterrupt() {
		return ErrTimeout
	}

	for id, task := range r.tasks {
		task(id)
	}

	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {

	case <-r.interrupt:
		// Stop receiving any further signals
		signal.Stop(r.interrupt)
		return true
		// Continue running as normal
	default:
		return false
	}
}
