package main

import (
	"fmt"
	"log"
	"net/http"

	"groupieTracker/functions"
)

func main() {
	// Route handlers
	http.HandleFunc("/", functions.HandlerIndex)
	http.HandleFunc("/static/", functions.HandleStatic)

	// Start server
	fmt.Println("Server running on: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
