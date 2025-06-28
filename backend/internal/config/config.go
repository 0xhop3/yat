package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port              string
	DatabaseURL       string
	Auth0Domain       string
	Auth0ClientID     string
	Auth0ClientSecret string
	JWTSecret         string
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, using default value: %v", err)
	}

	return &Config{
		Port:              getEnv("PORT", "9090"),
		DatabaseURL:       getEnv("DATABASE_URL", "postgres://0xskaper:061312@localhost:5234/yat_db?sslmode=disable"),
		Auth0Domain:       getEnv("AUTH0_DOMAIN", ""),
		Auth0ClientID:     getEnv("AUTH0_CLIENT_ID", ""),
		Auth0ClientSecret: getEnv("AUTH0_CLIENT_SECRET", ""),
		JWTSecret:         getEnv("JWT_SECRET", ""),
	}
}
