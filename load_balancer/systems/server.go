package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Response from backend on port %s\n", port)
	})

	log.Printf("Backend server listening on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
