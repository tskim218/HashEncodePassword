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

	// channel for a shutdown signal
	shutDown := make(chan string)

	// this creates the backend in-memory storage system
	db := datasource.NewInMemoryDB()

	// this creates the worker status system
	worker := datasource.WorkerStatus()

	// process stats
	stats := datasource.ProcStats()

	// this creates a new http.ServeMux, which is used to register handlers to execute in response to routes
	mux := http.NewServeMux()

	// get the password of a id
	mux.Handle("/hash/", handlers.GetPassword(db, worker))

	// create, encode, and set the value of a password
	mux.Handle("/hash", handlers.EncodePassword(db, worker, stats))

	// initiate shut down process
	mux.Handle("/shutdown", handlers.ShutDown(shutDown, worker))

	// get stats of the processes
	mux.Handle("/stats", handlers.GetStats(stats))

	log.Printf("serving on port 8080")

	srv := &http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf(":%s", "8080"),
	}

	go func() {
		err := srv.ListenAndServe()
		log.Fatal(err)
	}()

	// receive the shut down signal
	<-shutDown

	// Waiting until all current workers done
	if err := worker.WaitingForWorkersDone(); err != nil {
		log.Printf("error waiting current worker status %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("error shutting down server %s", err)
	} else {
		log.Println("Server gracefully stopped")
	}
}