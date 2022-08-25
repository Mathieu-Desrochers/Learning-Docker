package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		log.Print("GET /sum")

		json.NewEncoder(w).Encode(15)
	})

	log.Print("Listening on :8082...")

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal(err)
	}
}
