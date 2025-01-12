package service

import (
	"OSS-Matching-ServerSide/internal/model"
	"OSS-Matching-ServerSide/internal/repository"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService interface {
	Create(githubID, name, email string) (*model.User, error)
	Get(id uuid.UUID) (*model.User, error)
	Update(id uuid.UUID, name, email string) (*model.User, error)
	Delete(id uuid.UUID) error
}

type userService struct {
	db   *gorm.DB
	repo repository.UserRepository
}

func NewUserService(db *gorm.DB, repo repository.UserRepository) UserService {
	return &userService{
		db:   db,
		repo: repo,
	}
}

func (s *userService) Create(githubID string, name string, email string) (*model.User, error) {
	newUser := &model.User{
		ID:       uuid.New(),
		GithubID: githubID,
		Name:     name,
		Email:    email,
	}

	createdUser, err := s.repo.Create(s.db, newUser)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return createdUser, nil
}

func (s *userService) Get(id uuid.UUID) (*model.User, error) {
	user, err := s.repo.Get(s.db, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (s *userService) Update(id uuid.UUID, name, email string) (*model.User, error) {
	existingUser, err := s.repo.Get(s.db, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	existingUser.Name = name
	existingUser.Email = email

	if err := s.repo.Update(s.db, existingUser); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return existingUser, nil
}

func (s *userService) Delete(id uuid.UUID) error {
	if err := s.repo.Delete(s.db, id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
