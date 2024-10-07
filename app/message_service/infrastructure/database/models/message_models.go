package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	MessageID string         `gorm:"type:char(36);primaryKey" json:"message_id"`
	SenderID  string         `gorm:"type:char(36);not null" json:"sender_id"`
	Content   string         `gorm:"not null" json:"content"`
	MediaURL  *string        `gorm:"size:255"`
	SentAt    time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// for relationship with 'Conversation' Table
	ConvID       string       `gorm:"type:char(36);not null" json:"conv_id"` //foreign key
	Conversation Conversation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
