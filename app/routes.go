package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id   string
	Name string
}

func handleHelloWorld() http.HandlerFunc {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")

			user := User{Id: "123ABC", Name: "John Doe"}
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
