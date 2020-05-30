package handlers

import (
	"log"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"crypto/sha512"
    "encoding/base64"

	"github.com/tskim218/HashEncodePassword/datasource"
)

// PutKey returns an http.Handler that can set a value for the key registered by Gorilla
// mux as "key" in the path. It expects the value to be in the body of the PUT request
func PostPassword(db datasource.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if db.ShutDownStatus() == true {
			log.Printf("Can't accept your request due to shuting down")
			w.Write([]byte("Can't accept your request due to shuting down\n"))
			return
		}

		db.IncWorker()

		defer r.Body.Close()
		val, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error reading PUT body", http.StatusBadRequest)
			return
		}

		password := strings.Split(string(val), "=") 

		log.Printf("--- %s\n", password[1])

		hasher := sha512.New()
    	hasher.Write([]byte(password[1]))

    	sha := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

    	log.Printf("sha = %s\n", sha)

    	id := db.GetId()

    	go func() {
			//if val, err := db.Set(password[1], sha); err != nil {
			if err := db.Set(id, sha); err != nil {
				http.Error(w, "error setting value in DB", http.StatusInternalServerError)
				return
			}
		}()

		str := strconv.FormatUint(id, 10)

		log.Printf("why -%s\n", str)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(str+"\n"))
		//w.Write([]byte("working?"))
		log.Printf("returning")

		db.DecWorker()
	})
}