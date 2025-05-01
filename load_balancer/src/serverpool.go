package main

import (
	"sync/atomic"
)

// ServerPool is the struct for the server pool
type ServerPool struct {
	backends []*Backend
	current  uint64
}

var serverPool ServerPool

// methods in the server pool

// NextIndex checks and returns the index of the next server in the pool assuming there is another server in the pool
func (s *ServerPool) NextIndex() int {
	return int(atomic.AddUint64(&s.current, 1) % uint64(len(s.backends)))
}

// GetNextPeer gets the next server in the pool assuming there is another server in the pool
func (s *ServerPool) GetNextPeer() *Backend {
	next := s.NextIndex()
	l := len(s.backends) + next
	for i := next; i < l; i++ {
		idx := i % len(s.backends)
		if s.backends[idx].IsAlive() {
			if i != next {
				atomic.StoreUint64(&s.current, uint64(idx))
			}
			return s.backends[idx]
		}
	}
	return nil
}
