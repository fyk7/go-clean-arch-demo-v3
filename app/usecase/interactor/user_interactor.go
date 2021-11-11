package usecase

import (
	"context"
	"fmt"

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
	repo    repository.UserRepository
	service *service.UserService
}

func NewUserInteractor(repo repository.UserRepository, service *service.UserService) *userInteractor {
	return &userInteractor{
		repo:    repo,
		service: service,
	}
}

func (u *userInteractor) ListUser(c context.Context) ([]*model.User, error) {
	users, err := u.repo.FindAll(c)
	if err != nil {
		return nil, fmt.Errorf("failed to LidyUser in usecase: %w", err)
	}
	return users, nil
}

func (u *userInteractor) GetUserByEmail(c context.Context, email string) (*model.User, error) {
	user, err := u.repo.FindByEmail(c, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// TODO userを返すようにする
func (u *userInteractor) Create(c context.Context, email string) (*model.User, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	if err = u.service.Duplicated(c, email); err != nil {
		return nil, err
	}
	user := &model.User{
		ID:    uid.String(),
		Email: email,
	}
	user, err = u.repo.Save(c, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
