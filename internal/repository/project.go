package repository

import (
	"OSS-Matching-ServerSide/internal/model"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	Create(db *gorm.DB, project *model.Project) (*model.Project, error)
	Get(db *gorm.DB, id string) (*model.Project, error)
	Update(db *gorm.DB, project *model.Project) error
	Delete(db *gorm.DB, id string) error
}

type projectRepository struct{}

func NewProjectRepository() ProjectRepository {
	return &projectRepository{}
}

func (r *projectRepository) Create(db *gorm.DB, project *model.Project) (*model.Project, error) {
	if err := db.Create(project).Error; err != nil {
		return nil, err
	}
	return project, nil
}

func (r *projectRepository) Get(db *gorm.DB, id string) (*model.Project, error) {
	var project model.Project
	if err := db.First(&project, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *projectRepository) Update(db *gorm.DB, project *model.Project) error {
	return db.Save(project).Error
}

func (r *projectRepository) Delete(db *gorm.DB, id string) error {
	return db.Delete(&model.Project{}, "id = ?", id).Error
}
