package repository

import (
	"OSS-Matching-ServerSide/internal/model"

	"gorm.io/gorm"
)

type JobApplicationRepository interface {
	Create(db *gorm.DB, application *model.JobApplication) (*model.JobApplication, error)
	Get(db *gorm.DB, id string) (*model.JobApplication, error)
	Update(db *gorm.DB, application *model.JobApplication) error
	Delete(db *gorm.DB, id string) error
}

type jobApplicationRepository struct{}

func NewJobApplicationRepository() JobApplicationRepository {
	return &jobApplicationRepository{}
}

func (r *jobApplicationRepository) Create(db *gorm.DB, application *model.JobApplication) (*model.JobApplication, error) {
	if err := db.Create(application).Error; err != nil {
		return nil, err
	}
	return application, nil
}

func (r *jobApplicationRepository) Get(db *gorm.DB, id string) (*model.JobApplication, error) {
	var application model.JobApplication
	if err := db.First(&application, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &application, nil
}

func (r *jobApplicationRepository) Update(db *gorm.DB, application *model.JobApplication) error {
	return db.Save(application).Error
}

func (r *jobApplicationRepository) Delete(db *gorm.DB, id string) error {
	return db.Delete(&model.JobApplication{}, "id = ?", id).Error
}
