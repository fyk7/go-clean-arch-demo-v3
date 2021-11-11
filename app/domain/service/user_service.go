package service

import (
	"context"
	"fmt"
	"reflect"

	"github.com/fyk7/go-clean-arch-demo-v3/app/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Duplicated(c context.Context, email string) error {
	user, _ := s.repo.FindByEmail(c, email)
	if (user != nil) || !reflect.ValueOf(user).IsNil() {
		return fmt.Errorf("user %s is already exists", email)
	}
	// TODO errorの種別によってはerrorを返す必要がある
	return nil
}
