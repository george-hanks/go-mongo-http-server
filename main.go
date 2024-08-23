package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/george-hanks/go-mongo-http-server/app"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Missing PORT ENV variable")
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("Missing MONG_DB_URI ENV variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	usersCollection := client.Database("default").Collection("users")

	srv := app.NewServer(usersCollection)

	httpServer := &http.Server{
		Addr:    `:` + port,
		Handler: srv,
	}

	fmt.Println("Running Server On Port " + port)

	httpServer.ListenAndServe()
}
