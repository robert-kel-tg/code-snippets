package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type HandlerAdapter func(http.Handler) http.Handler

// Log every request specific fields
func LoggerRequest() HandlerAdapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			logEntryOf(r).Info("Request headers")

			h.ServeHTTP(w, r)

			log.Println("After log")
		})
	}
}

// Start and finish span for every request
func TracerRequest() HandlerAdapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("Before")
			
			h.ServeHTTP(w, r)
			
			log.Println("After")
		})
	}
}

func logEntryOf(r *http.Request) *log.Entry {
	return log.WithFields(log.Fields{
		"method": r.Method,
		"path":   r.URL.Path,
		"header": r.Header,
	})
}
