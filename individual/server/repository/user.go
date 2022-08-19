package repository

import "context"

type IUserRepository interface {
	Create(ctx context.Context)
}

type UserRepository struct {
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create(ctx context.Context) {}
