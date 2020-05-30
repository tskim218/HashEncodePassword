package handlers

import (
	"net/http"
)

func ShutDown(shutdown chan string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Shut Down\n"))
		shutdown <- "Shut Down"
	})
}