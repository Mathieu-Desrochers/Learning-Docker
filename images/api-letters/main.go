package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/letters", func(w http.ResponseWriter, r *http.Request) {
		log.Print("GET /letters")
		json.NewEncoder(w).Encode([]string{"A", "B", "C", "D"})
	})

	log.Print("Listening on :8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
