package main

import (
	"log"
	"net/http"
)

// main class, starts the backend servers
func main() {
	backends := []string{ // collect all the local servers into a list
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
	}

	// loop through list of servers, creating them into backend servers and adding to the pool
	for _, b := range backends {
		backend := NewBackend(b)
		serverPool.backends = append(serverPool.backends, backend)
	}

	// start the balancer
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(lb),
	} // traffic will flow to 8080, then redirected to 8081,8082 or 8083 depending on usage
	log.Println("Load Balancer started at 8080")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
