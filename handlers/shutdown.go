package handlers

import (
	"github.com/tskim218/HashEncodePassword/datasource"
	"log"
	"net/http"
)

// shut down signal is received by the route
func ShutDown(shutdown chan string, worker datasource.Worker) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := worker.ChangeShutDownStatus(); err != nil {
			log.Printf("error changing shut down status %s", err)
			return
		}

		w.Write([]byte("Shut Down\n"))
		shutdown <- "Shut Down"
	})
}