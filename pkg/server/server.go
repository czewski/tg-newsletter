package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func CreateServer() {
	http.HandleFunc("/new-message", messageHandler)

	fmt.Println("Starting Server")

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/new-message" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Bah")
}
