package app

import (
	"net/http"

	"github.com/george-hanks/go-mongo-http-server/app/handlers"

	"go.mongodb.org/mongo-driver/mongo"
)

func addRoutes(mux *http.ServeMux, usersCollection *mongo.Collection) {

	mux.Handle("GET /users", handlers.GetUsers(usersCollection))
	mux.Handle("GET /users/{id}", handlers.GetUser(usersCollection))
	mux.Handle("PUT /users", handlers.CreateUser(usersCollection))

	mux.Handle("/", http.NotFoundHandler())
}
