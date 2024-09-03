package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/george-hanks/go-mongo-http-server/app"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func run(logger *log.Logger) {
	port := os.Getenv("PORT")
	if port == "" {
		logger.Fatal("Missing PORT ENV variable")
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		logger.Fatal("Missing MONG_DB_URI ENV variable")
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

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalln(err)
	}

	usersCollection := client.Database("cyberpunk-red").Collection("users")

	srv := app.NewServer(usersCollection)

	httpServer := &http.Server{
		Addr:    `:` + port,
		Handler: srv,
	}

	logger.Println("Running Server On Port " + port)

	httpServer.ListenAndServe()
}

func main() {

	logger := log.New(os.Stdout, "[SERVER] ", log.LstdFlags)

	run(logger)

}
