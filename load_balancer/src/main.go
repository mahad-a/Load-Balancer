package main

import (
	"log"
	"net/http"
)

func main() {
	backends := []string{
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
	}

	for _, b := range backends {
		backend := NewBackend(b)
		serverPool.backends = append(serverPool.backends, backend)
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(lb),
	}
	log.Println("Load Balancer started at 8080")
	server.ListenAndServe()
}
