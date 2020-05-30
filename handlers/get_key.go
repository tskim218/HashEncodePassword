package handlers

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"net/http"
	"strings"

	"github.com/tskim218/HashEncodePassword/datasource"
)

// GetKey returns an http.Handler that can get a key registered by Gorilla mux
// as "key" in the path. It gets the value of the key from db
func GetKey(db datasource.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//key := r.URL.Query().Get(db.inc)
		//if key == "" {
		//	http.Error(w, "missing key name in query string", http.StatusBadRequest)
		//	return
		//}
		//val, err := db.Get(key)
		//if err == datasource.ErrNotFound {
		//	http.Error(w, "not found", http.StatusNotFound)
		//	return
		//} else if err != nil {
		//	http.Error(w, fmt.Sprintf("error getting value from database: %s", err), http.StatusInternalServerError)
		//	return
		//}

		db.IncWorker()

		parts := strings.Split(r.URL.String(), "/")
		if len(parts) != 3 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		//log.Printf("%s\n", r.URL.Path)

		u, err := strconv.ParseUint(parts[2], 10, 64)

		i := uint64(u)

		val, err := db.Get(i)
		if err == datasource.ErrNotFound {
			http.Error(w, "not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, fmt.Sprintf("error getting value from database: %s", err), http.StatusInternalServerError)
			return
		}
		for i := 0; i < 3; i++ {
			log.Printf("Still working...")
			time.Sleep(10*time.Second)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(val+"\n"))
		//w.Write([]byte("hello"))

		db.DecWorker()
	})
}