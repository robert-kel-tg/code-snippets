package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/satori/go.uuid"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	log "github.com/sirupsen/logrus"
)

var tokenAuth *jwtauth.JwtAuth

type User struct {
	ID       uuid.UUID
	Name     string
	Username string
}

func init() {

	// set logrus
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	user := User{
		ID:       uuid.NewV4(),
		Name:     "Robert",
		Username: "rke",
	}

	tokenAuth = jwtauth.New("HS512", []byte("secret"), nil)
	_, tokenString, _ := tokenAuth.Encode(jwtauth.Claims{"user": user})

	log.Warn("Debug: a sample jwt is %s ", tokenString)
}

func main() {

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())

			u, err := json.Marshal(claims)
			if err != nil {
				log.WithError(err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(u)
		})
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("It's Orders Services!")))

	})

	http.ListenAndServe(":8080", r)
}
