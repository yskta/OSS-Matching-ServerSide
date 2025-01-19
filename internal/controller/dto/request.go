package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	GithubID string `json:"github_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type UpdateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type CreateProjectRequest struct {
	GithubRepoID string `json:"github_repo_id" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Description  string `json:"description"`
}

type UpdateProjectRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type CreateProjectContributorRequest struct {
	ProjectID uuid.UUID `json:"project_id" validate:"required"`
	UserID    uuid.UUID `json:"user_id" validate:"required"`
	Role      string    `json:"role" validate:"required"`
}

type CreateJobPostingRequest struct {
	ProjectID   uuid.UUID `json:"project_id" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	Status      string    `json:"status" validate:"required"`
	Deadline    time.Time `json:"deadline"`
}

type CreateJobApplicationRequest struct {
	JobPostingID uuid.UUID `json:"job_posting_id" validate:"required"`
	UserID       uuid.UUID `json:"user_id" validate:"required"`
}

type CreateChatMessageRequest struct {
	JobApplicationID uuid.UUID `json:"job_application_id" validate:"required"`
	SenderID         uuid.UUID `json:"sender_id" validate:"required"`
	Content          string    `json:"content" validate:"required"`
}

type CreateRequiredSkillRequest struct {
	JobPostingID uuid.UUID `json:"job_posting_id" validate:"required"`
	Name         string    `json:"name" validate:"required"`
	Level        string    `json:"level"`
}

type CreateUserSkillRequest struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
	Name   string    `json:"name" validate:"required"`
	Level  string    `json:"level"`
}