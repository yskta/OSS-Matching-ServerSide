package service

import (
	"OSS-Matching-ServerSide/internal/model"
	"OSS-Matching-ServerSide/internal/repository"
	"database/sql"
	"fmt"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type ProjectContributorService interface {
	Create(projectID, userID uuid.UUID, role string) (*model.ProjectContributor, error)
	Get(projectID, userID uuid.UUID) (*model.ProjectContributor, error)
	Update(projectID, userID uuid.UUID, role string) (*model.ProjectContributor, error)
	Delete(projectID, userID uuid.UUID) error
}

type projectContributorService struct {
	db   *gorm.DB
	repo repository.ProjectContributorRepository
}

func NewProjectContributorService(db *gorm.DB, repo repository.ProjectContributorRepository) ProjectContributorService {
	return &projectContributorService{
		db:   db,
		repo: repo,
	}
}

func (s *projectContributorService) Create(projectID, userID uuid.UUID, role string) (*model.ProjectContributor, error) {
	newContributor := &model.ProjectContributor{
		ProjectID: projectID,
		UserID:    userID,
		Role:      role,
		CanManageJobPosting: sql.NullBool{
			Bool:  false, // デフォルトではfalse
			Valid: true,
		},
	}

	createdContributor, err := s.repo.Create(s.db, newContributor)
	if err != nil {
		return nil, fmt.Errorf("failed to create project contributor: %w", err)
	}

	return createdContributor, nil
}

func (s *projectContributorService) Get(projectID, userID string) (*model.ProjectContributor, error) {
	contributor, err := s.repo.Get(s.db, projectID, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get project contributor: %w", err)
	}

	return contributor, nil
}

func (s *projectContributorService) Update(projectID, userID, role string) (*model.ProjectContributor, error) {
	existingContributor, err := s.repo.Get(s.db, projectID, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get project contributor: %w", err)
	}

	existingContributor.Role = role

	if err := s.repo.Update(s.db, existingContributor); err != nil {
		return nil, fmt.Errorf("failed to update project contributor: %w", err)
	}

	return existingContributor, nil
}

func (s *projectContributorService) Delete(projectID, userID string) error {
	if err := s.repo.Delete(s.db, projectID, userID); err != nil {
		return fmt.Errorf("failed to delete project contributor: %w", err)
	}

	return nil
}
