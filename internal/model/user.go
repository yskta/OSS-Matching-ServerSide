package model

import "time"

type User struct {
	ID        string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	GithubID  string    `gorm:"type:varchar(255);unique;not null"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP"`
}
