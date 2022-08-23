package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/numbers", func(w http.ResponseWriter, r *http.Request) {
		log.Print("GET /numbers")
		json.NewEncoder(w).Encode([]int{1, 2, 3, 4, 5})
	})

	log.Print("Listening on :8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
