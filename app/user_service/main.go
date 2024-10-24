package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"user_service/infrastructure/database"
	"user_service/infrastructure/database/models"
	"user_service/infrastructure/repository"
	grpchandler "user_service/interface/grpc_handler"
	"user_service/usecase"

	"github.com/joho/godotenv"
	"github.com/ysmnababan/naniAppProto/pb"

	"google.golang.org/grpc"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Getenv("USER_PORT"))
}

func main() {
	db := database.Connect()
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Phonebook{})
	repo := &repository.Repo{DB: db}
	userUsecase := &usecase.UserUsecase{UserRepositoryI: repo}
	userHandler := &grpchandler.UserHandler{UserUsecaseI: userUsecase}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userHandler)

	listen, err := net.Listen("tcp", ":"+os.Getenv("USER_PORT"))
	if err != nil {
		log.Println(err)
	}

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}
