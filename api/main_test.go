package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/robertke/orders-service/pkg/timehandler"
)

func TestTimeEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/time", nil)

	if err != nil {
		t.Fatal(err)
	}

	// We need to record response
	rr := httptest.NewRecorder()
	th := timehandler.NewTimeHandler(time.RFC1123)
	th.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := time.Now().Format(time.RFC1123)

	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestPopulateContext(t *testing.T) {
	req, err := http.NewRequest("GET", "/time", nil)

	if err != nil {
		t.Fatal(err)
	}

	// We create test handler func
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if val, ok := r.Context().Value("app.req_id").(string); !ok {
			t.Errorf("app.req.id not in request context: got %q", val)
		}
	})

	rr := httptest.NewRecorder()
	handler := timehandler.RequestIDMiddleware(testHandler)
	handler.ServeHTTP(rr, req)
}
