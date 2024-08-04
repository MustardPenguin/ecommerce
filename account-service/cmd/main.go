package main

import (
	"account-service/cmd/util"
	"account-service/internal/api"
	"database/sql"
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
	db := util.InitConnection()

	// Start http server
	port := os.Getenv("PORT")
	api.StartServer(db, port)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing db")
		}
	}(db)
}
