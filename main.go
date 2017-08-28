package main

import "github.com/go-chi/chi"
import "net/http"

func main() {

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It's Orders Services!"))
	})
	http.ListenAndServe(":8080", r)
}
