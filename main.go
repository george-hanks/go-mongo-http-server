package main

import (
	"fmt"
	"main/app"
	"net/http"
)

func main() {

	srv := app.NewServer()

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: srv,
	}

	fmt.Println("Running Server On Port 8080")

	httpServer.ListenAndServe()
}
