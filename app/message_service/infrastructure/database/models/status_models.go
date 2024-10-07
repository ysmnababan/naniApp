package models

import "time"

type Status struct {
	StatusID  string    `gorm:"type:char(36);primaryKey" json:"status_id"`
	TargetID  string    `gorm:"type:char(36);not null" json:"target_id"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // enum for (sent, delivered, read)
	Info      string    `gorm:"size:50;not null" json:"info"`

	// for relationship with 'Message' Table
	MessageID string  `gorm:"type:char(36);not null" json:"message_id"` // foreign key field
	Message   Message `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
