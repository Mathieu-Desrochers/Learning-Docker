package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/numbers", func(w http.ResponseWriter, r *http.Request) {
		log.Print("GET /numbers")
		io.WriteString(w, "[1,2,3,4,5]\n")
	})

	log.Print("Listening on :8081...")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
