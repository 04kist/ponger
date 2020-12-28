package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	hostName, err := os.Hostname()
	if err != nil {
		log.Fatal("Somebody requested ping")
	}

	httpResponse := fmt.Sprintf("Pong from %s!\n", hostName)

	pong := func(w http.ResponseWriter, _ *http.Request) {
		log.Println("Http request")
		enableCors(&w)
		io.WriteString(w, httpResponse)
	}

	http.HandleFunc("/ping", pong)

	log.Println("Server started on port :8080")
	log.Println("Use curl <IP OR HOSTNAME>:8080/ping to get pong")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
