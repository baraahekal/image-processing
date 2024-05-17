package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hello", helloHandler)

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	if r.Method == "GET" {
		fmt.Fprintf(w, "Hello from the server!")
		return
	}

	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}
