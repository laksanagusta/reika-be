package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all application configuration
type Config struct {
	Server ServerConfig
	Gemini GeminiConfig
	CORS   CORSConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port string
}

// GeminiConfig holds Gemini API configuration
type GeminiConfig struct {
	APIKey string
}

// CORSConfig holds CORS configuration
type CORSConfig struct {
	AllowOrigins string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists (ignore error if file doesn't exist)
	_ = godotenv.Load()

	config := &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "5002"),
		},
		Gemini: GeminiConfig{
			APIKey: os.Getenv("GEMINI_API_KEY"),
		},
		CORS: CORSConfig{
			AllowOrigins: getEnv("CORS_ALLOW_ORIGINS", "http://localhost:3000"),
		},
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return config, nil
}

// Validate validates the configuration
func (c *Config) Validate() error {
	if c.Gemini.APIKey == "" {
		return errors.New("GEMINI_API_KEY environment variable is required")
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
