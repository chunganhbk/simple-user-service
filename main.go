package main

import (
	"fmt"
	"github.com/chunganhbk/simple-user-service/services"
	"github.com/chunganhbk/simple-user-service/repository"

	"github.com/chunganhbk/simple-user-service/database"
	pb "github.com/chunganhbk/simple-user-service/proto/auth"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	_, err := database.Storage.Connect()
	if err != nil {
		fmt.Println("Error: Failed to load database")
	}
	defer database.Storage.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()
	repo := &repository.UserRepository{ database.Storage.DB}

	tokenService := &services.TokenService{database.DBStorage}
	// Register Handler
	//pb.RegisterAuthHandler(service.Server(), &services.userService{ repo, tokenService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
