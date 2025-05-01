package main

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

// backend server structure and logic

// Backend is the struct for the backend server
type Backend struct {
	URL          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

// NewBackend creates a new backend server
func NewBackend(rawURL string) *Backend {
	parsedUrl, _ := url.Parse(rawURL)
	proxy := httputil.NewSingleHostReverseProxy(parsedUrl)

	return &Backend{
		URL:          parsedUrl,
		Alive:        true,
		ReverseProxy: proxy,
	}
}

// methods for the backend server

// SetAlive enters the mutex and changes the state of life of the server
func (b *Backend) SetAlive(alive bool) {
	b.mux.Lock()
	b.Alive = alive
	b.mux.Unlock()
}

// IsAlive checks the current state of life of the server
func (b *Backend) IsAlive() (alive bool) {
	b.mux.RLock()
	alive = b.Alive
	b.mux.RUnlock()
	return
}
