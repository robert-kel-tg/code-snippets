package mutex

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	value int
}

func Mutex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		counter := Counter{}

		for i := 0; i < 10; i++ {
			go func(number int) {
				counter.Lock()
				counter.value++
				defer counter.Unlock()
			}(i)
		}

		time.Sleep(time.Second)

		counter.Lock()
		defer counter.Unlock()

		w.Write([]byte(fmt.Sprintf("Count: %d", counter.value)))
	}
}
