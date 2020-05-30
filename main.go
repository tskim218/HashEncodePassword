package main

import (
	"context"
	"fmt"
	"time"

	"github.com/tskim218/HashEncodePassword/datasource"
	"github.com/tskim218/HashEncodePassword/handlers"
	"log"
	"net/http"
)

func main() {

	shutDown := make(chan string)

	// this creates the backend storage system
	db := datasource.NewInMemoryDB()

	// this creates a new http.ServeMux, which is used to register handlers to execute in response to routes
	mux := http.NewServeMux()

	// get the value of a key
	mux.Handle("/hash/", handlers.GetKey(db))

	// set the value of a key
	mux.Handle("/hash", handlers.PostPassword(db))

	mux.Handle("/shutdown", handlers.ShutDown(shutDown, db))

	log.Printf("serving on port 8080")

	srv := &http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf(":%s", "8080"),
	}

	go func() {
		err := srv.ListenAndServe()
		log.Fatal(err)
	}()

	<-shutDown

	// Waiting until all current workers done
	db.WaitingForWorkersDone()

	log.Printf("why not done")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("error shutting down server %s", err)
	} else {
		log.Println("Server gracefully stopped")
	}
}