package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleHelloWorld() http.HandlerFunc {

	type User struct {
		id   string
		name string
	}

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")

			user := User{id: "123ABC", name: "John Doe"}
			json.NewEncoder(w).Encode(user)

		},
	)
}

func loggerMiddleware(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request Made")
		h(w, r)
	})
}

func addRoutes(mux *http.ServeMux) {
	mux.Handle("POST /helloWorld", loggerMiddleware(handleHelloWorld()))

	mux.Handle("/", http.NotFoundHandler())
}

func NewServer() http.Handler {
	mux := http.NewServeMux()

	addRoutes(mux)

	var handler http.Handler = mux

	return handler
}

func main() {

	srv := NewServer()

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: srv,
	}

	fmt.Println("Running Server On Port 8080")

	httpServer.ListenAndServe()
}
