package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/strawHat121/diary-api/database"
	"github.com/strawHat121/diary-api/model"
)

func main() {
	loadEnv()
	loadDatabase()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
