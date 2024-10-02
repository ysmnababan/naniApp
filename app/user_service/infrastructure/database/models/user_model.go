package models

import (
	"time"

	"gorm.io/gorm"
)

// User model using UUID v7
type User struct {
	UserID      string         `gorm:"type:char(36);primaryKey" json:"user_id"` // Using UUID v7
	Username    string         `gorm:"size:50;not null" json:"username"`
	Email       string         `gorm:"size:320;uniqueIndex;not null" json:"email"`
	Password    string         `gorm:"size:255;not null" json:"password"`
	PhoneNumber string         `gorm:"size:20;uniqueIndex;not null" json:"phone_number"`
	PictureURL  *string        `gorm:"size:255" json:"picture_url"` //nullable field
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
