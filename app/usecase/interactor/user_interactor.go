package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/fyk7/go-clean-arch-demo-v3/app/domain/model"
	"github.com/fyk7/go-clean-arch-demo-v3/app/domain/repository"
	"github.com/fyk7/go-clean-arch-demo-v3/app/domain/service"
	"github.com/google/uuid"
)

type UserInteractor interface {
	ListUser(c context.Context) ([]*model.User, error)
	Create(c context.Context, email string) (*model.User, error)
	GetUserByEmail(c context.Context, email string) (*model.User, error)
}

type userInteractor struct {
	repo           repository.UserRepository
	service        *service.UserService
	contextTimeout time.Duration
}

func NewUserInteractor(repo repository.UserRepository, service *service.UserService, timeout time.Duration) *userInteractor {
	return &userInteractor{
		repo:           repo,
		service:        service,
		contextTimeout: timeout,
	}
}

func (u *userInteractor) ListUser(c context.Context) ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	users, err := u.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to LidyUser in usecase: %w", err)
	}
	return users, nil
}

func (u *userInteractor) GetUserByEmail(c context.Context, email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err := u.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// TODO userを返すようにする
func (u *userInteractor) Create(c context.Context, email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	if err = u.service.Duplicated(ctx, email); err != nil {
		return nil, err
	}
	user := &model.User{
		ID:    uid.String(),
		Email: email,
	}
	user, err = u.repo.Save(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
