package main

import (
	"log"
	"net/http"
)

func main() {
	// All URLs will be handled by this function
	// http.HandleFunc uses the DefaultServeMux
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	// Continue to process new requests until an error occurs
	err := http.ListenAndServe(":8090", nil)

	if err == nil {
		log.Println("Server started at http://localhost:8090")
	}

	log.Fatal(err)
}
