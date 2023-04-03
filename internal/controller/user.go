package controller

import (
	"user-management/domain"
	"user-management/repository/inmemory"
)

type UserController struct {
	// Explicit dependency that hides dependent logic
	store domain.UserStore
}

func NewUserController() *UserController {
	return &UserController{
		store: inmemory.NewUserStore(),
	}
}

func (c UserController) Create(u *domain.User) error {
	return c.store.Create(u)
}

func (c UserController) Get(userId string) (*domain.User, error) {
	return c.store.Get(userId)
}
