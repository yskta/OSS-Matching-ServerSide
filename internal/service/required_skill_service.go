package service

import (
	"OSS-Matching-ServerSide/internal/model"
	"OSS-Matching-ServerSide/internal/repository"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RequiredSkillService interface {
	Create(jobPostingID uuid.UUID, name, level string) (*model.RequiredSkill, error)
}

type requiredSkillService struct {
	db   *gorm.DB
	repo repository.RequiredSkillRepository
}

func NewRequiredSkillService(db *gorm.DB, repo repository.RequiredSkillRepository) RequiredSkillService {
	return &requiredSkillService{
		db:   db,
		repo: repo,
	}
}

func (s *requiredSkillService) Create(jobPostingID uuid.UUID, name, level string) (*model.RequiredSkill, error) {
	newRequiredSkill := &model.RequiredSkill{
		ID:           uuid.New(),
		JobPostingID: jobPostingID,
		Name:         name,
		Level: sql.NullString{
			String: level,
			Valid:  level != "", // 空文字列の場合はNULL
		},
	}
	createdRequiredSkill, err := s.repo.Create(s.db, newRequiredSkill)
	if err != nil {
		return nil, fmt.Errorf("failed to create chat message: %w", err)
	}

	return createdRequiredSkill, nil
}
