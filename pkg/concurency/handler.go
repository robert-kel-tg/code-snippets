package concurency

import (
	"fmt"
	"net/http"
	"sync"
)

// https://www.safaribooksonline.com/library/view/go-design-patterns/9781788390552/ch20s02.html
func Concurency(c string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var wait sync.WaitGroup

		routines := 5

		wait.Add(routines)

		for i := 0; i < routines; i++ {
			go func(writer http.ResponseWriter, name string, routineID int) {
				writer.Write([]byte(fmt.Sprintf("%s Bbbbbbbb with ID: %d\n", name, routineID)))
				// wait.Add(-1)
				wait.Done()
			}(w, c, i)
		}

		wait.Wait()
		// time.Sleep(time.Second)
	}
}
