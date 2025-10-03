package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("⚠️ .env file not found, using system environment variables")
	}
}

func GetKey(key string) string {
	return os.Getenv(key)
}
