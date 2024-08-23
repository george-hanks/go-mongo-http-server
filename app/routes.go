package app

import (
	"george-hanks/main/app/handlers"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func addRoutes(mux *http.ServeMux, usersCollection *mongo.Collection) {

	mux.Handle("GET /users", handlers.GetUsers(usersCollection))
	mux.Handle("GET /users/{id}", handlers.GetUser(usersCollection))
	mux.Handle("PUT /users", handlers.CreateUser(usersCollection))

	mux.Handle("/", http.NotFoundHandler())
}
