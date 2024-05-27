package main

import (
	"net/http"
	"fmt"
	"time"
)

const PORT = ":9090"
const CUSTOM_TOKEN = "Aklj12k3jkasdjflkasdjflkasjf13j"

func handler(w http.ResponseWriter, r *http.Request) {
	// Extract token from query parameters
	token := r.URL.Query().Get("token")

	if token == CUSTOM_TOKEN {
		fmt.Fprintf(w, "Request at %v\n", time.Now())
		for k, v := range r.Header {
			fmt.Fprintf(w, "%v: %v\n", k, v)
		}

		// Display the client IP address including possible proxy headers
		clientIP := r.RemoteAddr
		if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
			clientIP = forwardedFor
		} else if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
			clientIP = realIP
		}
		fmt.Fprintf(w, "Client IP: %v\n", clientIP)

		// Set and display response headers
		w.Header().Set("Content-Type", "text/plain")

	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(PORT, nil)
}
