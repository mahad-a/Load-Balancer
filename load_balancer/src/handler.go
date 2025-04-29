package main

import (
	"log"
	"net/http"
	"time"
)

func lb(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	peer := serverPool.GetNextPeer()
	if peer != nil {
		peer.ReverseProxy.ServeHTTP(w, r)
		duration := time.Since(start).Milliseconds()
		log.Printf("Request handled by %s, duration: %v milliseconds", peer.URL.String(), duration)
		return
	}
	http.Error(w, "Service not available", http.StatusServiceUnavailable)
}
