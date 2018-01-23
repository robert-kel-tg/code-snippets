package pubsub

import "testing"
import "sync"

type mockSubscriber struct {
	notifytestingFunc func(msg interface{})
	closeTestingFunc  func()
}

func (s *mockSubscriber) Close() {
	s.closeTestingFunc()
}

func (s *mockSubscriber) Notify(msg interface{}) error {
	s.notifytestingFunc(msg)
	return nil
}

func TestPublisher(t *testing.T) {
	msg := "Hello"

	p := NewPublisher()

	var wg sync.WaitGroup

	sub := &mockSubscriber{
		notifytestingFunc: func(msg interface{}) {
			defer wg.Done()

		},
		closeTestingFunc: func() {
			wg.Done()
		},
	}
}
