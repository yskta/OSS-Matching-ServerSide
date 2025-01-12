package service

import (
	"OSS-Matching-ServerSide/internal/model"
	"OSS-Matching-ServerSide/internal/repository"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserSkillService interface {
	Create(userID uuid.UUID, name, level string) (*model.UserSkill, error)
}

type userSkillService struct {
	db   *gorm.DB
	repo repository.UserSkillRepository
}

func NewUserSkillService(db *gorm.DB, repo repository.UserSkillRepository) UserSkillService {
	return &userSkillService{
		db:   db,
		repo: repo,
	}
}

func (s *userSkillService) Create(userID uuid.UUID, name, level string) (*model.UserSkill, error) {
	newUserSkill := &model.UserSkill{
		ID:     uuid.New(),
		UserID: userID,
		Name:   name,
		Level: sql.NullString{
			String: level,
			Valid:  level != "", // 空文字列の場合はNULL
		},
	}
	createdUserSkill, err := s.repo.Create(s.db, newUserSkill)
	if err != nil {
		return nil, fmt.Errorf("failed to create chat message: %w", err)
	}

	return createdUserSkill, nil
}
