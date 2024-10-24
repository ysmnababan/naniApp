package models

import (
	"time"
)

type Phonebook struct {
	PhonebookID string    `gorm:"type:char(36);primaryKey" json:"phonebook_id"` // Using UUID v7
	UserID      string    `gorm:"type:char(36);not null" json:"user_id"`
	ContactID   string    `gorm:"type:char(36);not null" json:"contact_id"`
	Nickname    string    `gorm:"size:320;not null" json:"nickname"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
