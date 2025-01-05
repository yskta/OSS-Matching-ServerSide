// internal/repository/required_skill.go
package repository

import (
	"OSS-Matching-ServerSide/internal/model"

	"gorm.io/gorm"
)

type RequiredSkillRepository interface {
	Create(db *gorm.DB, skill *model.RequiredSkill) (*model.RequiredSkill, error)
	Get(db *gorm.DB, id string) (*model.RequiredSkill, error)
	Update(db *gorm.DB, skill *model.RequiredSkill) error
	Delete(db *gorm.DB, id string) error
}

type requiredSkillRepository struct{}

func NewRequiredSkillRepository() RequiredSkillRepository {
	return &requiredSkillRepository{}
}

func (r *requiredSkillRepository) Create(db *gorm.DB, skill *model.RequiredSkill) (*model.RequiredSkill, error) {
	if err := db.Create(skill).Error; err != nil {
		return nil, err
	}
	return skill, nil
}

func (r *requiredSkillRepository) Get(db *gorm.DB, id string) (*model.RequiredSkill, error) {
	var skill model.RequiredSkill
	if err := db.First(&skill, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &skill, nil
}

func (r *requiredSkillRepository) Update(db *gorm.DB, skill *model.RequiredSkill) error {
	return db.Save(skill).Error
}

func (r *requiredSkillRepository) Delete(db *gorm.DB, id string) error {
	return db.Delete(&model.RequiredSkill{}, "id = ?", id).Error
}
