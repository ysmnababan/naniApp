package main

import (
	"log"
	"message_service/infrastructure/database"
	"message_service/infrastructure/repository"
	"message_service/usecase"
	"os"

	"github.com/joho/godotenv"
	"github.com/ysmnababan/naniAppProto/pb"
	"google.golang.org/grpc"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	conn, err := grpc.Dial("localhost:"+os.Getenv("USER_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	db := database.Connect()
	messageRepo := &repository.Repo{DB: db}
	messageUsecase := &usecase.MessageUsecase{MessageRepoI: messageRepo, UserClient: client}
	_ = messageUsecase
}
