package usecase

import (
	"context"
	"errors"

	"github.com/go-dep/models"
	"github.com/go-dep/repo"
)

type UserUsecase struct {
	UserRepo repo.UserRepository
}

type UserService interface {
	GetUser(ctx context.Context, name string) (models.User, error)
	AddUser(ctx context.Context, name string) error
	DeleteUser(ctx context.Context, name string) error
	UpdateUser(ctx context.Context, oldName string, newName string) error
}

func (s UserUsecase) GetUser(ctx context.Context, name string) (models.User, error) {

	if len(name) < 3 {
		return models.User{}, errors.New("Short name")
	}

	user, err := s.UserRepo.GetUser(ctx, name)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s UserUsecase) AddUser(ctx context.Context, name string) error {
	if len(name) < 3 {
		return errors.New("Short name")
	}

	err := s.UserRepo.AddUser(ctx, name)

	if err != nil {
		return err
	}

	return nil
}

func (s UserUsecase) DeleteUser(ctx context.Context, name string) error {
	err := s.UserRepo.DeleteUser(ctx, name)

	if err != nil {
		return err
	}

	return nil
}

func (s UserUsecase) UpdateUser(ctx context.Context, oldName string, newName string) error {
	if len(newName) < 3 {
		return errors.New("Short name")
	}

	err := s.UserRepo.UpdateUser(ctx, oldName, newName)

	if err != nil {
		return err
	}

	return nil
}
