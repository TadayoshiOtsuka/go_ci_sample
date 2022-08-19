package repository

import (
	"context"

	"github.com/TadayoshiOtsuka/go_test_sample/app/individual/server/entity"
	"github.com/TadayoshiOtsuka/go_test_sample/app/pkg/rdb"
)

type IUserRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
}

type UserRepository struct {
	db *rdb.SQLHandler
}

func NewUserRepository(db *rdb.SQLHandler) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	u, err := r.db.Client.User.
		Create().
		SetName(user.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity.NewUser(u.ID, u.Name)
}
