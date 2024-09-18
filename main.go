package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/george-hanks/go-mongo-http-server/app"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func run(ctx context.Context) error {

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Missing PORT ENV variable")
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("Missing MONG_DB_URI ENV variable")
	}

	fmt.Println("Attempting To Connect To MongoDB")
	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}

	usersCollection := client.Database("cyberpunk-red").Collection("users")

	srv := app.NewServer(usersCollection)

	httpServer := &http.Server{
		Addr:    `:` + port,
		Handler: srv,
	}

	go func() {
		fmt.Println("Running Server On Port " + port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		os.Exit(1)
	}
}
