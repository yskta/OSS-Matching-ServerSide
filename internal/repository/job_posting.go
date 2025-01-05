package repository

import (
	"OSS-Matching-ServerSide/internal/model"

	"gorm.io/gorm"
)

type JobPostingRepository interface {
	Create(db *gorm.DB, posting *model.JobPosting) (*model.JobPosting, error)
	Get(db *gorm.DB, id string) (*model.JobPosting, error)
	Update(db *gorm.DB, posting *model.JobPosting) error
	Delete(db *gorm.DB, id string) error
}

type jobPostingRepository struct{}

func NewJobPostingRepository() JobPostingRepository {
	return &jobPostingRepository{}
}

func (r *jobPostingRepository) Create(db *gorm.DB, posting *model.JobPosting) (*model.JobPosting, error) {
	if err := db.Create(posting).Error; err != nil {
		return nil, err
	}
	return posting, nil
}
func (r *jobPostingRepository) Get(db *gorm.DB, id string) (*model.JobPosting, error) {
	var posting model.JobPosting
	if err := db.First(&posting, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &posting, nil
}

func (r *jobPostingRepository) Update(db *gorm.DB, posting *model.JobPosting) error {
	return db.Save(posting).Error
}

func (r *jobPostingRepository) Delete(db *gorm.DB, id string) error {
	return db.Delete(&model.JobPosting{}, "id = ?", id).Error
}
