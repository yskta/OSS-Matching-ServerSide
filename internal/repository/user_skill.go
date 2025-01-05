// internal/repository/user_skill.go
package repository

import (
	"OSS-Matching-ServerSide/internal/model"

	"gorm.io/gorm"
)

type UserSkillRepository interface {
	Create(db *gorm.DB, skill *model.UserSkill) (*model.UserSkill, error)
	Get(db *gorm.DB, id string) (*model.UserSkill, error)
	Update(db *gorm.DB, skill *model.UserSkill) error
	Delete(db *gorm.DB, id string) error
}

type userSkillRepository struct{}

func NewUserSkillRepository() UserSkillRepository {
	return &userSkillRepository{}
}

func (r *userSkillRepository) Create(db *gorm.DB, skill *model.UserSkill) (*model.UserSkill, error) {
	if err := db.Create(skill).Error; err != nil {
		return nil, err
	}
	return skill, nil
}

func (r *userSkillRepository) Get(db *gorm.DB, id string) (*model.UserSkill, error) {
	var skill model.UserSkill
	if err := db.First(&skill, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &skill, nil
}

func (r *userSkillRepository) Update(db *gorm.DB, skill *model.UserSkill) error {
	return db.Save(skill).Error
}

func (r *userSkillRepository) Delete(db *gorm.DB, id string) error {
	return db.Delete(&model.UserSkill{}, "id = ?", id).Error
}
