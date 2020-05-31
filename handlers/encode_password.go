package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"crypto/sha512"
	"encoding/base64"

	"github.com/tskim218/HashEncodePassword/datasource"
)

// Returns an http.Handler that can set a password for the key registered by mux.
// It expects the password to be in the data of the POST request
func EncodePassword(db datasource.DB, worker datasource.Worker) http.Handler {
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

		defer r.Body.Close()
		val, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error reading PUT body", http.StatusBadRequest)
			return
		}

		// register the current worker
		worker.IncWorker()

		log.Printf("----- request encoding password ------\n\n")

		// parse the data
		password := strings.Split(string(val), "=")
		if len(password) != 2 {
			log.Printf("error wrong format of data, %s - %s", string(val), err)
			worker.DecWorker()
			return
		}

		log.Printf("password is received: %s\n", password[1])

		// encode the password
		hasher := sha512.New()
    	hasher.Write([]byte(password[1]))

    	encode := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

    	log.Printf("Encoded password = %s\n", encode)

    	// get id for the password
    	id := db.GetId()

    	go func() {
			if err := db.Set(id, encode); err != nil {
				http.Error(w, "error setting value in DB", http.StatusInternalServerError)
				worker.DecWorker()
				return
			}
		}()

		w.WriteHeader(http.StatusOK)

		str := strconv.FormatUint(id, 10)
		log.Printf("server returns id (%s)\n", str)

		// return to the client
		w.Write([]byte(str+"\n"))

		// unregister the current worker
		worker.DecWorker()
	})
}