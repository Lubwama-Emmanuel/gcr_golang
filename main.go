package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	port := 3000
	fmt.Println("Server is running on port", port)

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
