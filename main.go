package main

import (
	"log"
	"net/http"
)

func main() {
	setupApi()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupApi() {

	manager := NewManager()
	// serve static frontend files
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	// suppress automatic browser favicon requests to avoid 404 noise
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	http.HandleFunc("/ws", manager.serverWS)
}
