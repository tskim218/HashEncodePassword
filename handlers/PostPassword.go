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
		// key := r.URL.Query().Get("password")
		// if key == "" {
		// 	http.Error(w, "missing key name in path", http.StatusBadRequest)
		// 	return
		// }
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
    	// fmt.Printf("hmac512:\t%s\n", base64.StdEncoding.EncodeToString(hmac512.Sum(nil)))
    	// sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
    	sha := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

    	log.Printf("sha = %s\n", sha)

    	id := db.GetId()
		//w.Write([]byte(string(id)))

    	go func() {
			//if val, err := db.Set(password[1], sha); err != nil {
			if err := db.Set(id, sha); err != nil {
				http.Error(w, "error setting value in DB", http.StatusInternalServerError)
				return
			}
		}()

		//go func() {
		//	db.Set(password[1], sha)
		//}()


		// reqBody, err := ioutil.ReadAll(r.Body)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 		http.Error(w, "error reading PUT body", http.StatusBadRequest)
		// 		return
		// 	}
		// defer r.Body.Close()

		// log.Printf("%s\n", reqBody)
		// password := strings.Split(string(reqBody), ",") 
		// log.Printf("%s\n", password)


		str := strconv.FormatUint(id, 10)

		log.Printf("why -%s\n", str)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(str+"\n"))
		//w.Write([]byte("working?"))
		log.Printf("returning")
	})
}

// func (ms *MapServer) storee(bv []byte) {
//     hasher := sha512.New()
//     hasher.Write(bv)
//     sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
// }
