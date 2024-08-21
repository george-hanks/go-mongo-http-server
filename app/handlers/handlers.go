package handlers

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Id   string
	Name string
}

func HandleHelloWorld() http.HandlerFunc {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")

			user := User{Id: "123ABC", Name: "John Doe"}
			json.NewEncoder(w).Encode(user)

		},
	)
}
