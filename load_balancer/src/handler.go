package main

import (
	"log"
	"net/http"
	"time"
)

// lb is the load balancer, handles incoming requests and decides which backend server should process each request
func lb(w http.ResponseWriter, r *http.Request) {
	start := time.Now() // timer to track the duration of the request handling

	// go through the pool, using round-robin to ensure equal/similar workload
	peer := serverPool.GetNextPeer()
	if peer != nil { // if a valid backend server is available
		peer.ReverseProxy.ServeHTTP(w, r)
		duration := time.Since(start).Milliseconds()
		log.Printf("Request handled by %s, duration: %v milliseconds", peer.URL.String(), duration)
		return
	}
	http.Error(w, "Service not available", http.StatusServiceUnavailable)
}
