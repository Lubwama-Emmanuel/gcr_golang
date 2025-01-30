package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World! welcome to cloud run, yeah it works well")
}

// healthCheckHandler responds to health check requests
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "healthy")
}

func main() {
	port := 8080
	fmt.Println("Server is running on port", port)

	http.HandleFunc("/test", handler)
	http.HandleFunc("/health", healthCheckHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
