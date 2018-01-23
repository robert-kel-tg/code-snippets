package pubsub

import "testing"
import "sync"
import "strings"
import "fmt"

func TestStdoutPrinter(t *testing.T) {

}

func TestWriter(t *testing.T) {
	msg := "Hello"

	var wg sync.WaitGroup
	wg.Add(1)

	sub := NewWriterSubscriber(0, nil)

	stdoutPrinter := sub.(*writerSubscriber)
	stdoutPrinter.Writer = &mockWriter{
		testingFunc: func(res string) {
			if !strings.Contains(res, msg) {
				t.Fatal(fmt.Errorf("Incorrect string: %s", res))
			}
			wg.Done()
		},
	}

	err := sub.Notify(msg)
	if err != nil {
		wg.Done()
		t.Error(err)
	}
	wg.Wait()
	sub.Close()
}

type mockWriter struct {
	testingFunc func(string)
}

func (mw *mockWriter) Write(p []byte) (n int, err error) {
	mw.testingFunc(string(p))
	return len(p), nil
}
