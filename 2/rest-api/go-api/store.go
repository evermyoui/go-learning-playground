package main

import (
	"sync"
	"time"
)

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

func (s *UserStore) List() []User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]User, 0, len(s.users))
	for _, u := range s.users {
		users = append(users, u)
	}
	return users
}

func (s *UserStore) Create(name, email string) User {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.counter++
	user := User{
		ID:        s.counter,
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}

	s.users[user.ID] = user
	return user
}
