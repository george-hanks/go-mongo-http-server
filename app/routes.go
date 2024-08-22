package app

import (
	"george-hanks/main/app/handlers"
	"george-hanks/main/app/middleware"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func addRoutes(mux *http.ServeMux, usersCollection *mongo.Collection) {
	mux.Handle("POST /helloWorld", middleware.LoggerMiddleware(handlers.HandleHelloWorld()))

	mux.Handle("GET /users", handlers.GetUsers(usersCollection))

	mux.Handle("/", http.NotFoundHandler())
}
