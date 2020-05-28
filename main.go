package main

import (
	"context"
	"fmt"
	"github.com/tskim218/HashEncodePassword/datasource"
	"github.com/tskim218/HashEncodePassword/handlers"
	"log"
	"net/http"
	"time"
)

func main() {

	messages := make(chan string)

	// this creates the backend storage system
	db := datasource.NewInMemoryDB()
	// this creates a new http.ServeMux, which is used to register handlers to execute in response to routes
	mux := http.NewServeMux()
	// get the value of a key
	mux.Handle("/get", handlers.GetKey(db))
	// set the value of a key
	mux.Handle("/hash", handlers.PostPassword(db))

	log.Printf("serving on port 8080")

	srv := &http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf(":%s", "8080"),
	}

	// http.ListenAndServe takes in an http.Handler as its second parameter.
	// since ServeMux implements a ServeHTTP function, it is also an http.Handler,
	// so we can pass it here.
	//err := http.ListenAndServe(":8080", mux)
	go func() {
		err := srv.ListenAndServe()
		log.Fatal(err)
	}()

	go func() { messages <- "ping" }()

	time.Sleep(10*time.Second)

	<-messages
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("error shutting down server %s", err)
	} else {
		log.Println("Server gracefully stopped")
	}
}