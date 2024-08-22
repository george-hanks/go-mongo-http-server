package main

import (
	"fmt"
	"main/app"
	"net/http"
	"os"
)

func main() {

	port, isSet := os.LookupEnv("PORT")
	if !isSet {
		panic("Missing PORT ENV variable")
	}

	srv := app.NewServer()

	httpServer := &http.Server{
		Addr:    `:` + port,
		Handler: srv,
	}

	fmt.Println("Running Server On Port " + port)

	httpServer.ListenAndServe()
}
