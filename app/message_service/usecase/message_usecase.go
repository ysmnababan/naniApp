package usecase

import (
	"context"
	"message_service/domain"
	"message_service/helper"
	"message_service/infrastructure/repository"

	"github.com/samborkent/uuidv7"
	"github.com/ysmnababan/naniAppProto/pb"
)

type MessageUsecase struct {
	repository.MessageRepoI
	UserClient pb.UserServiceClient
}

type MessageUsecaseI interface {
	CreateGroupChat(group *domain.Conversation) (*domain.Conversation, error)
	CreatePrivateChat(group *domain.Conversation) (*domain.Conversation, error)
	AddNewMember(user_id, conv_id string) (*domain.Conversation, error)

	SaveMessage(conv_id string, msg *domain.Message) (*domain.Message, error)
	GetConversationMessage(conv_id string) (*domain.Conversation, error)
	UnsentMessage(msg *domain.Message) error
	GetMemberList(conv_id string) ([]*domain.Participant, error)
	ReadByList(message_id string) ([]*domain.Participant, error)
}

func (u *MessageUsecase) CreateGroupChat(group *domain.Conversation) (*domain.Conversation, error) {
	if group.ConvName == "" || group.ConvType != "group" {
		return nil, helper.ErrParam
	}

	if len(group.Participants) == 0 {
		return nil, helper.ErrParam
	}

	// ensure each member is exist
	for i := range group.Participants {
		_, err := u.UserClient.GetUser(context.Background(),
		&pb.GetUserReq{UserId: group.Participants[i].UserID})
		if err != nil {
			return nil, err
		}
		
		// generate participant id
		group.Participants[i].ParticipantID = generateUUIDv7()
	}

	// creates uuid
	group.ConvID = generateUUIDv7()

	err := u.MessageRepoI.CreateConversation(group)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (u *MessageUsecase)CreatePrivateChat(group *domain.Conversation) (*domain.Conversation, error){
	if group.ConvType != "private" {
		return nil, helper.ErrParam
	}

	// private chat contains 2 member
	if len(group.Participants) != 2 {
		return nil, helper.ErrParam
	}


	// creates uuid
	group.ConvID = generateUUIDv7()


	// 
	return nil, nil
}


func generateUUIDv7() string {
	// creates uuid
	uuidV7 := uuidv7.New()
	return uuidV7.String()
}
