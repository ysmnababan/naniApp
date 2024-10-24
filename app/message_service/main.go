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
	app_env := os.Getenv("APP_ENV")

	var userClient pb.UserServiceClient
	if app_env == "development" {
		conn, err := grpc.Dial(os.Getenv("GRPC_USER_SERVICE_URL"), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		userClient = pb.NewUserServiceClient(conn)
	}

	db := database.Connect()
	messageRepo := &repository.Repo{DB: db}
	messageUsecase := &usecase.MessageUsecase{MessageRepoI: messageRepo, UserClient: userClient}
	_ = messageUsecase

}
