package user

import (
	"context"
	"cproject/internal/models"
	"errors"
	"fmt"
	"log"
)

type contextKey string

var userKey contextKey = "user"

type Service interface {
	IsRegisteredUser(user *models.User) (*models.User, error)
	RegisterUser(user *models.User) (*models.User, error)
}

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{repo}
}

func (s *service) IsRegisteredUser(user *models.User) (*models.User, error) {
	user, err := s.repo.isRegisteredUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (s *service) RegisterUser(user *models.User) (*models.User, error) {

	registerUser, err := s.repo.registerUser(user)
	if err != nil {
		return nil, err
	}

	return registerUser, nil

}

func (s *service) updateLikes(userResource *UserResource) error {

	err := s.repo.updateLikes(userResource)
	if err != nil {
		log.Println("error adding to user_resource", err.Error())
		return err
	}

	return nil

}

func withUser(ctx context.Context, user *models.User) context.Context {
	fmt.Printf("%v user is heree?", user)
	newContext := context.WithValue(ctx, userKey, user)
	fmt.Println("New context", newContext)
	return newContext
}

func GetUser(ctx context.Context) (*models.User, error) {
	user, ok := ctx.Value(userKey).(*models.User)
	if !ok {
		log.Println("user not found", user)
		return nil, errors.New("user not found")
	}

	return user, nil

}
