package breaker

import (
	"fmt"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
)

func HystrixBreaker() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		hystrix.ConfigureCommand("demoCommand", hystrix.CommandConfig{
			Timeout:               20,
			MaxConcurrentRequests: 1,
			ErrorPercentThreshold: 1,
		})

		hystrix.Go("demoCommand", func() error {

			// time.Sleep(2 * time.Second)
			fmt.Println("Service operational")
			//get smth from service
			return nil
		}, func(err error) error {
			fmt.Println("Service down")
			return nil
		})
	}
}
