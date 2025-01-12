package repository

import (
	"OSS-Matching-ServerSide/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectContributorRepository interface {
	Create(db *gorm.DB, contributor *model.ProjectContributor) (*model.ProjectContributor, error)
	Get(db *gorm.DB, projectID, userID uuid.UUID) (*model.ProjectContributor, error)
	Update(db *gorm.DB, contributor *model.ProjectContributor) error
	Delete(db *gorm.DB, projectID, userID uuid.UUID) error
}

type projectContributorRepository struct{}

func NewProjectContributorRepository() ProjectContributorRepository {
	return &projectContributorRepository{}
}

func (r *projectContributorRepository) Create(db *gorm.DB, contributor *model.ProjectContributor) (*model.ProjectContributor, error) {
	if err := db.Create(contributor).Error; err != nil {
		return nil, err
	}
	return contributor, nil
}

func (r *projectContributorRepository) Get(db *gorm.DB, projectID, userID uuid.UUID) (*model.ProjectContributor, error) {
	var contributor model.ProjectContributor
	if err := db.First(&contributor, "project_id = ? AND user_id = ?", projectID, userID).Error; err != nil {
		return nil, err
	}
	return &contributor, nil
}

func (r *projectContributorRepository) Update(db *gorm.DB, contributor *model.ProjectContributor) error {
	return db.Save(contributor).Error
}

func (r *projectContributorRepository) Delete(db *gorm.DB, projectID, userID uuid.UUID) error {
	return db.Delete(&model.ProjectContributor{}, "project_id = ? AND user_id = ?", projectID, userID).Error
}
