package inmemory

import (
	"fmt"
	"user-management/domain"
)

var us *userStore = nil

type userStore struct {
	// An in-memory store with a map
	// Use User.ID as the key of map
	store map[string]*domain.User
}

// NewUserStore creates a new user store if it doesn't already exist.
func NewUserStore() *userStore {
	if us != nil {
		return us
	}

	us = &userStore{store: make(map[string]*domain.User)}
	return us
}

func (c userStore) Create(u *domain.User) error {
	if _, ok := c.store[u.Id]; ok {
		return fmt.Errorf("user already present with id %s", u.Id)
	}

	c.store[u.Id] = u
	return nil
}

func (c userStore) Get(userId string) (*domain.User, error) {
	user, ok := c.store[userId]
	if !ok {
		return nil, fmt.Errorf("no user with id %s", userId)
	}
	return user, nil
}
