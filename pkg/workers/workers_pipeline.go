package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

type Request struct {
	Data    interface{}
	Handler RequestHandler
}

type RequestHandler func(interface{})

func NewStringRequest(s string, id int, wg *sync.WaitGroup) Request {
	return Request{
		Data: "Hello",
		Handler: RequestHandler(func(i interface{}) {
			defer wg.Done()
			s, ok := i.(string)

			if !ok {
				log.Fatal("Invalid cast to string")
			}
			fmt.Println(s)
		}),
	}
}

func main() {
	bufferSize := 100

	var d Dispatcher
	d = NewDispatcher(bufferSize)

	workers := 3
	for i := 0; i < workers; i++ {
		var w WorkerLauncher
		w = NewPreffixSuffixWorker(
			i,
			fmt.Sprintf("WorkerID: %d -> ", i),
			" World",
		)
		d.LaunchWorker(i, w)
	}

	requests := 10
	var wg sync.WaitGroup
	wg.Add(requests)

	for i := 0; i < requests; i++ {
		req := NewStringRequest("(Msg_id: %d) -> Hello", i, &wg)
		d.MakeRequest(req)
	}
	d.Stop()
	wg.Wait()
}

// worker.go
type WorkerLauncher interface {
	LaunchWorker(i int, in chan Request)
}

type PreffixSuffixWorker struct {
	id      int
	prefixS string
	suffixS string
}

func NewPreffixSuffixWorker(id int, prefixS string, suffixS string) *PreffixSuffixWorker {
	return &PreffixSuffixWorker{
		id:      id,
		prefixS: prefixS,
		suffixS: suffixS,
	}
}

func (w *PreffixSuffixWorker) LaunchWorker(i int, in chan Request) {
	w.prefix(w.append(w.uppercase(in)))
}

func (w *PreffixSuffixWorker) uppercase(in <-chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)

			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = strings.ToUpper(s)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PreffixSuffixWorker) append(in <-chan Request) <-chan Request {
	out := make(chan Request)
	go func() {
		for msg := range in {
			uppercaseString, ok := msg.Data.(string)

			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = fmt.Sprintf("%s%s", uppercaseString, w.suffixS)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PreffixSuffixWorker) prefix(in <-chan Request) {
	go func() {
		for msg := range in {
			uppercasedStringWithSuffix, ok := msg.Data.(string)

			if !ok {
				msg.Handler(nil)
				continue
			}

			msg.Handler(fmt.Sprintf("%s%s", w.prefixS, uppercasedStringWithSuffix))
		}
	}()
}

// dispacher.go
// Launching workers in parallel and handling all the possible incoming channels
type Dispatcher interface {
	LaunchWorker(i int, w WorkerLauncher)
	MakeRequest(Request)
	Stop()
}

type dispatcher struct {
	inChannel chan Request
}

func NewDispatcher(b int) Dispatcher {
	return &dispatcher{
		inChannel: make(chan Request, b),
	}
}

func (d *dispatcher) LaunchWorker(id int, w WorkerLauncher) {
	w.LaunchWorker(id, d.inChannel)
}

func (d *dispatcher) MakeRequest(r Request) {
	select {
	case d.inChannel <- r:
	case <-time.After(time.Second * 5):
		return
	}
}

func (d *dispatcher) Stop() {
	close(d.inChannel)
}
