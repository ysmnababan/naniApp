package main

import (
	"user_service/infrastructure/database"
	"user_service/infrastructure/repository"
)

func main() {
	db := database.Connect()

	Repo := &repository.Repo{DB: db}
	
	Repo.DB.AutoMigrate()
}
