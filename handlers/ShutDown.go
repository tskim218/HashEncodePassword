package handlers

import (
	"github.com/tskim218/HashEncodePassword/datasource"
	"net/http"
)

func ShutDown(shutdown chan string, db datasource.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Shut Down\n"))
		db.ChangeShutDownStatus()
		shutdown <- "Shut Down"
	})
}