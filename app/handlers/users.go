package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
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

			var users []User
			if err = cursor.All(context.TODO(), &users); err != nil {
				json.NewEncoder(w).Encode([]string{})
				return
			}

			json.NewEncoder(w).Encode(users)

		},
	)
}

func GetUser(usersCollection *mongo.Collection) http.HandlerFunc {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			userId := r.PathValue("id")

			if userId == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			var user User
			err := usersCollection.FindOne(context.TODO(), bson.M{"_id": userId}).Decode(&user)
			if err != nil {
				json.NewEncoder(w).Encode([]string{})
				return
			}

			json.NewEncoder(w).Encode(user)

		},
	)
}

func CreateUser(usersCollection *mongo.Collection) http.HandlerFunc {

	type CreateUserBody struct {
		Name string `bson:"name"`
	}

	return http.HandlerFunc(

		func(w http.ResponseWriter, r *http.Request) {

			var requestBody CreateUserBody
			err := json.NewDecoder(r.Body).Decode(&requestBody)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			insertResult, err := usersCollection.InsertOne(context.TODO(), requestBody)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(insertResult.InsertedID)

		},
	)
}
