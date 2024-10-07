package repository

import (
	"message_service/domain"
	"message_service/infrastructure/database/models"

	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

type MessageRepoI interface {
	CreateConversation(conv *domain.Conversation) error
}

func (r *Repo) CreateConversation(conv *domain.Conversation) error {
	member := []models.Participant{}
	for _, participant := range conv.Participants {
		member = append(member, models.Participant{
			ParticipantID: participant.ParticipantID,
			UserID:        participant.UserID,
			ConvID:        conv.ConvID,
		})
	}

	convModel := models.Conversation{
		ConvID:       conv.ConvID,
		ConvName:     conv.ConvName,
		ConvType:     conv.ConvType,
		Participants: member,
	}
	
	res := r.DB.Create(&convModel)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
