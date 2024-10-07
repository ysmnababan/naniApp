package models

import (
	"time"

	"gorm.io/gorm"
)

type Conversation struct {
	ConvID       string         `gorm:"type:char(36);primaryKey" json:"conv_id"`
	ConvName     string         `gorm:"size:255;" json:"conv_name"`
	ConvType     string         `gorm:"size:20;not null" json:"conv_type"` // enum for (PRIVATE, GROUP)
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// relationship for the 'Participant' Table and 'Message' Table
	Participants []Participant  `gorm:"foreignKey;ConvID"`
	Messages     []Message      `gorm:"foreignKey;ConvID"`
}
