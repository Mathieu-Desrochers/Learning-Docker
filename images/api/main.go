package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		log.Print("GET /sum")

		_, err := http.Get("http://database:8080/numbers")
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(15)
	})

	log.Print("Listening on :8081...")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
