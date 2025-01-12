package service

import (
	"OSS-Matching-ServerSide/internal/model"
	"OSS-Matching-ServerSide/internal/repository"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChatMessageService interface {
	Create(jobApplicationID, senderID uuid.UUID, content string) (*model.ChatMessage, error)
}

type chatMessageService struct {
	db   *gorm.DB
	repo repository.ChatMessageRepository
}

func NewChatMessageService(db *gorm.DB, repo repository.ChatMessageRepository) ChatMessageService {
	return &chatMessageService{
		db:   db,
		repo: repo,
	}
}

func (s *chatMessageService) Create(jobApplicationID, senderID uuid.UUID, content string) (*model.ChatMessage, error) {
	newChatMessage := &model.ChatMessage{
		ID:               uuid.New(),
		JobApplicationID: jobApplicationID,
		SenderID:         senderID,
		Content:          content,
	}
	createdChatMessage, err := s.repo.Create(s.db, newChatMessage)
	if err != nil {
		return nil, fmt.Errorf("failed to create chat message: %w", err)
	}

	return createdChatMessage, nil
}
