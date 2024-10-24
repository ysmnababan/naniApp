package grpchandler

import (
	"context"
	"user_service/domain"
	"user_service/helper"
	"user_service/usecase"

	"github.com/ysmnababan/naniAppProto/pb"
)

type UserHandler struct {
	usecase.UserUsecaseI
}

func (h *UserHandler) GetContact(ctx context.Context, in *pb.GetContactReq) (*pb.GetContactResp, error) {
	if in.UserId == "" {
		return nil, helper.ErrInvalidId
	}

	res, err := h.UserUsecaseI.GetContacts(in.UserId)
	if err != nil {
		return nil, err
	}

	out := []*pb.Phonebook{}
	for i := range res {
		phonebook := &pb.Phonebook{
			PhonebookId: res[i].PhonebookID,
			UserId:      res[i].UserID,
			ContactId:   res[i].ContactID,
			Nickname:    res[i].Nickname,
		}
		out = append(out, phonebook)
	}
	return &pb.GetContactResp{Phonebooks: out}, nil
}

func (h *UserHandler) CreateContact(ctx context.Context, in *pb.CreateContactReq) (*pb.CreateContactResp, error) {
	contact := &domain.Phonebook{
		UserID: in.UserId,
	}
	res, err := h.UserUsecaseI.CreateNewContact(in.PhoneNumber, contact)
	if err != nil {
		return nil, err
	}

	phonebook := &pb.Phonebook{
		PhonebookId: res.PhonebookID,
		UserId:      res.UserID,
		ContactId:   res.ContactID,
		Nickname:    res.Nickname,
	}
	return &pb.CreateContactResp{Contact: phonebook}, nil
}

func (h *UserHandler) EditNickname(ctx context.Context, in *pb.EditNicknameReq) (*pb.EditNicknameResp, error) {
	res, err := h.UserUsecaseI.EditNickname(in.Nickname, in.PhonebookId)
	if err != nil {
		return nil, err
	}

	phonebook := &pb.Phonebook{
		PhonebookId: res.PhonebookID,
		UserId:      res.UserID,
		ContactId:   res.ContactID,
		Nickname:    res.Nickname,
	}
	return &pb.EditNicknameResp{Contact: phonebook}, nil
}

func (h *UserHandler) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	// validation
	if in.Email == "" || in.Password == "" {
		return nil, helper.ErrCredential
	}

	token, err := h.UserUsecaseI.Login(in.Email, in.Password)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResp{Token: token}, nil
}

func (h *UserHandler) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// validation
	if in.Email == "" || in.Password == "" || in.PhoneNumber == "" || in.Username == "" {
		return nil, helper.ErrParam
	}

	user := domain.User{
		Username:    in.Username,
		Email:       in.Email,
		Password:    in.Password,
		PhoneNumber: in.PhoneNumber,
		Picture_URL: &in.PictureUrl,
	}
	err := h.UserUsecaseI.Register(&user)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResp{Email: in.Email, Username: in.Username}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, in *pb.GetUserReq) (*pb.GetUserResp, error) {
	if in.UserId == "" {
		return nil, helper.ErrInvalidId
	}
	user, err := h.UserUsecaseI.GetUserData(in.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResp{
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		PictureUrl:  *user.Picture_URL,
	}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, in *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
	if in.UserId == "" {
		return nil, helper.ErrInvalidId
	}

	user := domain.User{
		UserID:      in.UserId,
		Username:    in.Username,
		Picture_URL: &in.PictureUrl,
	}

	updatedUser, err := h.UserUsecaseI.UpdateUserData(&user)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserResp{
		Username:    updatedUser.Username,
		Email:       updatedUser.Email,
		PhoneNumber: updatedUser.PhoneNumber,
		PictureUrl:  *updatedUser.Picture_URL,
	}, nil
}
