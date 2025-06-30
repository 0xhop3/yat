package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port              string
	DatabaseURL       string
	Auth0ID           string
	Auth0Domain       string
	Auth0Audience     string
	Auth0ClientID     string
	Auth0ClientSecret string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, using default values: %v", err)
	}

	return &Config{
		Port:          getEnv("PORT", ""),
		DatabaseURL:   getEnv("DATABASE_URL", ""),
		Auth0ID:       getEnv("AUTH0_ID", ""),
		Auth0Domain:   getEnv("AUTH0_DOMAIN", ""),
		Auth0Audience: getEnv("AUTH0_AUDIENCE", ""),
		Auth0ClientID: getEnv("AUTH0_CLIENT_ID", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
