package concurency

import (
	"errors"
	"io"
	"log"
	"sync"
)

// https://www.safaribooksonline.com/library/view/go-in-action/9781617291784/kindle_split_015.html
// Pool manages set of resources that can be
// shared safely by multiple go routines
type ResourcePool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrPoolClosed = errors.New("Pool has been closed.")

func New(fn func() (io.Closer, error), size uint) (*ResourcePool, error) {

	if size <= 0 {
		return nil, errors.New("Pool size value is too small.")
	}

	return &ResourcePool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Retrieve resourse from the Pool
func (p *ResourcePool) Retrieve() (io.Closer, error) {

	select {
	case r, ok := <-p.resources:
		log.Println("Pull:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Pull:", "New Resource")
		return p.factory()
	}
}

// Releases resource onto the Pool
func (p *ResourcePool) Release(r io.Closer) {

	p.m.Lock()
	defer p.m.Unlock()

	// If the pool is closed, discard the resource
	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("Release:", "In Queue")

	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

func (p *ResourcePool) Close() {

	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	// Close the channel before we drain it
	// Need to close overwise will have a deadlock
	close(p.resources)

	for r := range p.resources {
		r.Close()
	}
}
