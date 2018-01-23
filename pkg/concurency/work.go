package concurency

import (
	"sync"
)

// https://www.safaribooksonline.com/library/view/go-in-action/9781617291784/kindle_split_015.html
type Worker interface {
	Task()
}

type Pool interface {
	Run(Worker)
	Shutdown()
}

type pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func NewPool(maxGourotines int) *pool {
	p := pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGourotines)

	for i := 0; i < maxGourotines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

func (p *pool) Run(w Worker) {
	p.work <- w
}

func (p *pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
