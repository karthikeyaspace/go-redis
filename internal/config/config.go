package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	RedisAddr     string
	RedisUsername string
	RedisPassword string
	RedisDB       int
}

func NewConfig() *Config {
	loadEnv()
	return &Config{
		Port:          ":8080",
		RedisAddr:     os.Getenv("REDIS_ADDR"),
		RedisUsername: os.Getenv("REDIS_USERNAME"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       0,
	}
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	log.Print("Loaded .env file")
}
