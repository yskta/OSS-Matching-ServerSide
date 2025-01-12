package service

import (
	"OSS-Matching-ServerSide/internal/model"
	"OSS-Matching-ServerSide/internal/repository"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectService interface {
	Create(githubRepoID, name, description string) (*model.Project, error)
	Get(id uuid.UUID) (*model.Project, error)
	Update(id uuid.UUID, name, description string) (*model.Project, error)
	Delete(id uuid.UUID) error
}

type projectService struct {
	db   *gorm.DB
	repo repository.ProjectRepository
}

func NewProjectService(db *gorm.DB, repo repository.ProjectRepository) ProjectService {
	return &projectService{
		db:   db,
		repo: repo,
	}
}

func (s *projectService) Create(githubRepoID, name, description string) (*model.Project, error) {
	newProject := &model.Project{
		ID:           uuid.New(), // UUIDを生成
		GithubRepoID: githubRepoID,
		Name:         name,
		Description: sql.NullString{
			String: description,
			Valid:  description != "", // 空文字列の場合はNULL
		},
		IsActive: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
	}

	createdProject, err := s.repo.Create(s.db, newProject)

	if err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}

	return createdProject, nil
}

func (s *projectService) Get(id uuid.UUID) (*model.Project, error) {
	project, err := s.repo.Get(s.db, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get project: %w", err)
	}

	return project, nil
}

func (s *projectService) Update(id uuid.UUID, name, description string) (*model.Project, error) {
	existingProject, err := s.repo.Get(s.db, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get project: %w", err)
	}

	existingProject.Name = name
	existingProject.Description = sql.NullString{
		String: description,
		Valid:  description != "", // 空文字列の場合はNULL
	}

	if err := s.repo.Update(s.db, existingProject); err != nil {
		return nil, fmt.Errorf("failed to update project: %w", err)
	}

	return existingProject, nil
}

func (s *projectService) Delete(id uuid.UUID) error {
	if err := s.repo.Delete(s.db, id); err != nil {
		return fmt.Errorf("failed to delete project: %w", err)
	}

	return nil
}
