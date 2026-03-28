package main

import "sync"

// holds users
type UserStore struct {
	mu      sync.RWMutex // protects the map from concurrent
	users   map[int]User
	counter int //id
}

func NewUserStore() *UserStore {
	return &UserStore{
		users: make(map[int]User),
	}
}

func (s *UserStore) List() User {
	s.mu.RLock()

	return
}
