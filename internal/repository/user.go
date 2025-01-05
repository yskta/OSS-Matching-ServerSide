package repository

import (
	"OSS-Matching-ServerSide/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(db *gorm.DB, user *model.User) (*model.User, error)
	Get(db *gorm.DB, id string) (*model.User, error)
	Update(db *gorm.DB, user *model.User) error
	Delete(db *gorm.DB, id string) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(db *gorm.DB, user *model.User) (*model.User, error) {
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Get(db *gorm.DB, id string) (*model.User, error) {
	var user model.User
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(db *gorm.DB, user *model.User) error {
	return db.Save(user).Error
}

func (r *userRepository) Delete(db *gorm.DB, id string) error {
	return db.Delete(&model.User{}, "id = ?", id).Error
}
