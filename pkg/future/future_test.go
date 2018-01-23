package future

import "testing"
import "sync"
import "time"

func TestStringOrError_Execute(t *testing.T) {
	future := MaybeString{}

	t.Run("Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		// timeout
		time.Sleep(time.Second)
		t.Log("Timeout!")
		t.Fail()
		wg.Done()
		// timeout

		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})

		future.Execute(func(string, error) {

		})
		wg.Wait()
	})

	t.Run("Error result", func(t *testing.T) {

	})
}

// func timeout(t *testing.T, wg *sync.WaitGroup) {
// 	time.Sleep(time.Second)
// 	t.Log("Timeout!")

// 	t.Fail()
// 	wg.Done()
// }
