package main

import (
	"github.com/chunganhbk/simple-user-service/database"
	"github.com/chunganhbk/simple-user-service/repository"
	"github.com/chunganhbk/simple-user-service/services"
	"github.com/joho/godotenv"

	pb "github.com/chunganhbk/simple-user-service/proto/auth"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.CreateDBConnection()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	defer db.Close()
	var repo = &repository.UserRepository{DB: db}
	var tokenService = &services.TokenService{Repo: repo} // New Service
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()
	// Register Handler
	pb.RegisterAuthHandler(service.Server(), &services.UserService{ Repo: repo, TokenService: tokenService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

