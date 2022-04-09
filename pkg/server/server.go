package server

import (
	"fmt"
	"net/http"
)

func CreateServer() {
	http.HandleFunc("/new-message", messageHandler)

	fmt.Println("Starting Server")

	if err := http.ListenAndServe(":443", nil); err != nil {
		fmt.Println(err)
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
