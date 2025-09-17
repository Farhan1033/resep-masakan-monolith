package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("ENV file not found!")
	}
}

func GetKey(key string) string {
	return os.Getenv(key)
}
