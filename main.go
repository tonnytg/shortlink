
// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func jsonError(msg string) []byte {
	error := struct {
		Message string `json:"message"`
	}{
		msg,
	}
	r, err := json.Marshal(error)
	if err != nil {
		return []byte(err.Error())
	}
	return r
}

type Article struct {
    Title string `json:"name"`
	Body string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		article := Article{}

		err := json.NewDecoder(r.Body).Decode(&article)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		fmt.Fprintf(w, "Msg: %s!\n", article)
    }
}