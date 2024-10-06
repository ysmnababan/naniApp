package main

import (
	"log"
	"net"
	"user_service/infrastructure/database"
	"user_service/infrastructure/database/models"
	"user_service/infrastructure/repository"
	grpchandler "user_service/interface/grpc_handler"
	"github.com/ysmnababan/naniAppProto/pb"
	"user_service/usecase"

	"google.golang.org/grpc"
)

func main() {
	db := database.Connect()
	db.AutoMigrate(models.User{})
	repo := &repository.Repo{DB: db}
	userUsecase := &usecase.UserUsecase{UserRepositoryI: repo}
	userHandler := &grpchandler.UserHandler{UserUsecaseI: userUsecase}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userHandler)
	listen, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Println(err)
	}

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}
