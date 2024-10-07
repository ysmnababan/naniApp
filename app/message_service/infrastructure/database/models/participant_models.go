package models

import "time"

type Participant struct {
	ParticipantID string       `gorm:"type:char(36);primaryKey" json:"participant_id"`
	UserID        string       `gorm:"type:char(36);not null" json:"user_id"`
	JoinAt        time.Time    `gorm:"autoCreateTime"`

	// for relationship with 'Conversation' Table
	ConvID        string       `gorm:"type:char(36);not null" json:"conv_id"` // foreign key field
	Conversation  Conversation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
