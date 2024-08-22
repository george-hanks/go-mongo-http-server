package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id   string
	Name string
}

func HandleHelloWorld() http.HandlerFunc {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)

			user := User{Id: "123ABC", Name: "John Doe"}
			json.NewEncoder(w).Encode(user)

		},
	)
}

func GetUsers(usersCollection *mongo.Collection) http.HandlerFunc {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			cursor, err := usersCollection.Find(context.TODO(), bson.D{})
			if err != nil {
				json.NewEncoder(w).Encode([]string{})
				return
			}
			defer cursor.Close(context.TODO())

			var results []User
			if err = cursor.All(context.TODO(), &results); err != nil {
				json.NewEncoder(w).Encode([]string{})
				return
			}

			json.NewEncoder(w).Encode(results)

		},
	)
}
