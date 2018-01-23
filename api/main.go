package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/robertke/orders-service/pkg/concurency"
	"github.com/robertke/orders-service/pkg/middleware"

	"github.com/robertke/orders-service/pkg/tracing"
	"github.com/robertke/orders-service/pkg/usertype"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

// It’s common to set the log configuration inside of init()
// so the log package can be used immediately when the program starts
func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func main() {

	r := chi.NewRouter()

	r.Use(middleware.LoggerRequest())

	r.Get("/home", tracing.HomeHandler)
	r.Get("/db", tracing.DbHandler)
	r.Get("/async", tracing.ServiceHandler)
	r.Get("/service", tracing.ServiceHandler)

	r.Get("/user_type", usertype.UserTypeHandler)
	r.Get("/concurency", concurency.Concurency("test"))
	r.Get("/pool", func(w http.ResponseWriter, r *http.Request) {

	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("It's Orders Services!")))
	})

	http.ListenAndServe(":8080", r)
}
