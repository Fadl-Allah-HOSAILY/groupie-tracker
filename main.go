package main

import (
	"fmt"
	"log"
	"net/http"

	fn "groupieTracker/functions"
)


func main() {
	// Route handlers
	http.HandleFunc("/", fn.HandlerIndex)
	http.HandleFunc("/static/", fn.HandlerStatic)
	http.HandleFunc("/artist/", fn.HandlerArtist)

	// Start server
	fmt.Println("Server running on: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
