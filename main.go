package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id   string
	Name string
}

func handleHelloWorld() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")

			u := User{Id: "US123", Name: "John Doe"}
			json.NewEncoder(w).Encode(u)

		},
	)
}

func addRoutes(mux *http.ServeMux) {
	mux.Handle("POST /helloWorld", handleHelloWorld())
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

	fmt.Printf("Running Server On Port 8080")

	httpServer.ListenAndServe()
}
