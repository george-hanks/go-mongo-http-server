package app

import (
	"main/app/handlers"
	"main/app/middleware"
	"net/http"
)

func addRoutes(mux *http.ServeMux) {
	mux.Handle("POST /helloWorld", middleware.LoggerMiddleware(handlers.HandleHelloWorld()))

	mux.Handle("/", http.NotFoundHandler())
}
