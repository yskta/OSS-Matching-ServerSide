package dto

import (
	"time"
)

type UserResponse struct {
	ID        string    `json:"id"`
	GithubID  string    `json:"github_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProjectResponse struct {
	ID           string    `json:"id"`
	GithubRepoID string    `json:"github_repo_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description,omitempty"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ProjectContributorResponse struct {
	ProjectID           string    `json:"project_id"`
	UserID              string    `json:"user_id"`
	Role                string    `json:"role"`
	CanManageJobPosting bool      `json:"can_manage_job_posting"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type JobPostingResponse struct {
	ID          string    `json:"id"`
	ProjectID   string    `json:"project_id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status"`
	Deadline    time.Time `json:"deadline"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type JobApplicationResponse struct {
	ID           string    `json:"id"`
	JobPostingID string    `json:"job_posting_id"`
	UserID       string    `json:"user_id"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ChatMessageResponse struct {
	ID               string    `json:"id"`
	JobApplicationID string    `json:"job_application_id"`
	SenderID         string    `json:"sender_id"`
	Content          string    `json:"content"`
	IsRead           bool      `json:"is_read"`
	CreatedAt        time.Time `json:"created_at"`
}

type RequiredSkillResponse struct {
	ID           string    `json:"id"`
	JobPostingID string    `json:"job_posting_id"`
	Name         string    `json:"name"`
	Level        string    `json:"level,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserSkillResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Level     string    `json:"level,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
