package service

import (
	"OSS-Matching-ServerSide/internal/model"
	"OSS-Matching-ServerSide/internal/repository"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobApplicationService interface {
	Create(jobPostingID, userID uuid.UUID) (*model.JobApplication, error)
}

type jobApplicationService struct {
	db   *gorm.DB
	repo repository.JobApplicationRepository
}

func NewJobApplicationService(db *gorm.DB, repo repository.JobApplicationRepository) JobApplicationService {
	return &jobApplicationService{
		db:   db,
		repo: repo,
	}
}

func (s *jobApplicationService) Create(jobPostingID, userID uuid.UUID) (*model.JobApplication, error) {
	newJobApplication := &model.JobApplication{
		ID:           uuid.New(), // UUIDを生成
		JobPostingID: jobPostingID,
		UserID:       userID,
		Status:       model.JobApplicationStatusProgress,
	}

	createdJobApplication, err := s.repo.Create(s.db, newJobApplication)

	if err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}

	return createdJobApplication, nil
}
