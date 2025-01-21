package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World! welcome to cloud run, yeah it works well")
}

func main() {
	port := 8080
	fmt.Println("Server is running on port", port)

	http.HandleFunc("/test", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
