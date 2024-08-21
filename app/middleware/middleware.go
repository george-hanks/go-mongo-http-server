package middleware

import (
	"fmt"
	"net/http"
)

func LoggerMiddleware(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request Made")
		h(w, r)
	})
}
