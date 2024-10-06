package grpchandler

import (
	"context"
	"user_service/domain"
	"user_service/helper"
	"user_service/interface/grpc_handler/pb"
	"user_service/usecase"
)

type UserHandler struct {
	usecase.UserUsecaseI
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
