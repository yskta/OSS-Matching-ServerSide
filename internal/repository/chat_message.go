// internal/repository/chat_message.go
package repository

import (
	"OSS-Matching-ServerSide/internal/model"

	"gorm.io/gorm"
)

type ChatMessageRepository interface {
	Create(db *gorm.DB, message *model.ChatMessage) (*model.ChatMessage, error)
	Get(db *gorm.DB, id string) (*model.ChatMessage, error)
	Update(db *gorm.DB, message *model.ChatMessage) error
	Delete(db *gorm.DB, id string) error
}

type chatMessageRepository struct{}

func NewChatMessageRepository() ChatMessageRepository {
	return &chatMessageRepository{}
}

func (r *chatMessageRepository) Create(db *gorm.DB, message *model.ChatMessage) (*model.ChatMessage, error) {
	if err := db.Create(message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

func (r *chatMessageRepository) Get(db *gorm.DB, id string) (*model.ChatMessage, error) {
	var message model.ChatMessage
	if err := db.First(&message, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *chatMessageRepository) Update(db *gorm.DB, message *model.ChatMessage) error {
	return db.Save(message).Error
}

func (r *chatMessageRepository) Delete(db *gorm.DB, id string) error {
	return db.Delete(&model.ChatMessage{}, "id = ?", id).Error
}
