package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(key string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env file: %s", err)
	}
	return os.Getenv(key)
}
