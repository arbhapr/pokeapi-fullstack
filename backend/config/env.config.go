package config

import (
	"fmt"
	"log"
	"os"
	"poke-go/model"
	"strconv"

	"github.com/joho/godotenv"
)

// getEnv fetches environment variables or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("%s not set, using default: %s", key, defaultValue)
		return defaultValue
	}
	return value
}

// LoadEnv loads environment variables and returns the configuration struct
func LoadEnv() model.Config {
	// Only load .env file in development or if explicitly needed
	if os.Getenv("ENVIRONMENT") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Printf("Error loading .env file: %v", err)
		}
	}

	// Fetch environment variables with defaults
	environment := getEnv("ENVIRONMENT", "development")
	apiVersion := getEnv("API_VERSION", "v1")
	baseURL := getEnv("BASE_URL", "http://localhost")
	sourceURL := getEnv("SOURCE_URL", "http://pokeapi.co/api/v2")
	port := getEnv("PORT", "3000")

	// Convert LIMIT_PER_PAGE to an integer, with default
	limitPerPageStr := getEnv("LIMIT_PER_PAGE", "10")
	limitPerPage, err := strconv.Atoi(limitPerPageStr)
	if err != nil {
		log.Printf("Invalid LIMIT_PER_PAGE value, using default: 10")
		limitPerPage = 10
	}

	baseURL = fmt.Sprintf("%s:%s/api/%s", baseURL, port, apiVersion)
	log.Printf("%s", baseURL)

	// Construct and return the Config struct
	return model.Config{
		APIVersion:   apiVersion,
		BaseURL:      baseURL,
		SourceURL:    sourceURL,
		Port:         port,
		LimitPerPage: limitPerPage,
		Environment:  environment,
	}
}
