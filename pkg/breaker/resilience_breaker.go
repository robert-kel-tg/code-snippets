package breaker

import (
	"fmt"
	"net/http"
	"time"

	"gopkg.in/eapache/go-resiliency.v1/breaker"
)

// errorThreshold - is the number of times a request can fail before the circuit opens
// successThreshold - is the number of times that we need a successful request in the half-open state before we move back to open
// timeout - is the time that the circuit will stay in the open state before changing to half-open
func CirquitBreaker() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b := breaker.New(3, 1, 5*time.Second)

		for {
			result := b.Run(func() error {
				// Call service here

				fmt.Println("Getting service endpoint")

				time.Sleep(2 * time.Second)

				return fmt.Errorf("Timeout")
			})

			switch result {
			case nil:
				// success
			case breaker.ErrBreakerOpen:
				// our function wasn't run because the breaker was open
				fmt.Println("Breaker open")
			default:
				fmt.Println(result)
			}

			time.Sleep(500 * time.Millisecond)
		}
	}
}
