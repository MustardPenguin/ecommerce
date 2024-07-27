package main

import (
	"account-service/cmd/util"
	"account-service/internal/api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	// Load environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initiate database connection
	util.InitConnection()

	// Start http server
	port := os.Getenv("PORT")
	api.StartServer(port)
}
