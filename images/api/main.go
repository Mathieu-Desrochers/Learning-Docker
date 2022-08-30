package main

import (
	"io"
	"log"
	"net"
	"net/http"
)

func main() {
	var localIP net.IP

	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			localIP = ipnet.IP
		}
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello from "+localIP.String()+"\n")
	})

	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		log.Print("GET /sum")

		_, err := http.Get("http://database:8081/numbers")
		if err != nil {
		   log.Fatalln(err)
		}

		io.WriteString(w, "15\n")
	})

	log.Print("Listening on :8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
