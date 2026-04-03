package main

import (
	"fmt"
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

func (s *UserStore) GetById(id int) (User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, ok := s.users[id]
	if !ok {
		return User{}, fmt.Errorf("user with id %d not found", id)
	}
	return user, nil
}

func (s *UserStore) Update(id int, name, email string) (User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, ok := s.users[id]

	if !ok {
		return User{}, fmt.Errorf("user with id %d not found", id)
	}

	if name != "" {
		user.Name = name
	}

	if email != "" {
		user.Email = email
	}

	s.users[user.ID] = user

	return user, nil
}

func (s *UserStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.users[id]; !ok {
		return fmt.Errorf("user with id %d not found", id)
	}

	delete(s.users, id)

	return nil
}
