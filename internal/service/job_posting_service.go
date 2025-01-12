package service

import (
	"OSS-Matching-ServerSide/internal/model"
	"OSS-Matching-ServerSide/internal/repository"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobPostingService interface {
	Create(projectID uuid.UUID, title, description, status string, deadline time.Time) (*model.JobPosting, error)
}

type jobPostingService struct {
	db   *gorm.DB
	repo repository.JobPostingRepository
}

func NewJobPostingService(db *gorm.DB, repo repository.JobPostingRepository) JobPostingService {
	return &jobPostingService{
		db:   db,
		repo: repo,
	}
}

func ParseJobPostingStatus(status string) (model.JobPostingStatus, error) {
	switch status {
	case "draft":
		return model.JobPostingStatusDraft, nil
	case "open":
		return model.JobPostingStatusOpen, nil
	case "closed":
		return model.JobPostingStatusClosed, nil
	case "cancelled":
		return model.JobPostingStatusCancelled, nil
	default:
		return 0, fmt.Errorf("invalid status: %s", status)
	}
}

func (s *jobPostingService) Create(projectID uuid.UUID, title, description, status string, deadline time.Time) (*model.JobPosting, error) {
	// ステータスの変換
	jobStatus, err := ParseJobPostingStatus(status)
	if err != nil {
		return nil, err
	}

	// deadlineをsql.NullTimeに変換
	nullDeadline := sql.NullTime{
		Time:  deadline,
		Valid: !deadline.IsZero(), // deadlineが空でない場合はTrue
	}

	newJobPosting := &model.JobPosting{
		ID:        uuid.New(), // UUIDを生成
		ProjectID: projectID,
		Title:     title,
		Description: sql.NullString{
			String: description,
			Valid:  description != "", // 空文字列の場合はNULL
		},
		Status:   jobStatus,
		Deadline: nullDeadline,
	}

	createdJobPosting, err := s.repo.Create(s.db, newJobPosting)

	if err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}

	return createdJobPosting, nil
}
