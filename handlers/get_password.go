package handlers

import (
	"fmt"
	"github.com/tskim218/HashEncodePassword/datasource"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// GetPassword returns an http.Handler that can get an id.
// It gets the value of the key from db
func GetPassword(db datasource.DB, worker datasource.Worker) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		status, err := worker.ShutDownStatus()
		if err != nil {
			log.Printf("error shutting down status %s", err)
			return
		} else if status {
			log.Printf("Can't accept your request due to shutting down")
			w.Write([]byte("Can't accept your request due to shutting down\n"))
			return
		}

		// parse the URL path
		parts := strings.Split(r.URL.String(), "/")
		if len(parts) != 3 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// register the current worker
		worker.IncWorker()

		log.Printf("----- request getting password ------\n\n")

		id, err := strconv.ParseUint(parts[2], 10, 64)

		val, err := db.Get(id)
		if err == datasource.ErrNotFound {
			http.Error(w, "id not found", http.StatusNotFound)
			worker.DecWorker()
			return
		} else if err != nil {
			http.Error(w, fmt.Sprintf("error getting value from database: %s", err), http.StatusInternalServerError)
			worker.DecWorker()
			return
		}

		// simulate shutting down process waits until all processes are done.
		//for i := 0; i < 3; i++ {
		//	log.Printf("Still working...")
		//	time.Sleep(10*time.Second)
		//}

		log.Printf("Getting id: %d, password: %s", id, val)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(val+"\n"))

		// unregister the current worker
		worker.DecWorker()
	})
}