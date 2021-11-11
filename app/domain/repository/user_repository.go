package repository

import (
	"context"

	"github.com/fyk7/go-clean-arch-demo-v3/app/domain/model"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	Save(ctx context.Context, user *model.User) (*model.User, error)
}
