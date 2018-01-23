package timehandler

import (
	"context"
	"net/http"
	"time"
)

//https://elithrar.github.io/article/testing-http-handlers-go/
type TimeHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type timeHandler struct {
	format string
}

func NewTimeHandler(time string) TimeHandler {
	return &timeHandler{
		format: time,
	}
}

func (t *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(t.format)
	w.Write([]byte(tm))
}

func RequestIDMiddleware(h http.Handler) TimeHandler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "app.req_id", "12345")
		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
