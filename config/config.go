package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all application configuration
type Config struct {
	Server       ServerConfig
	Gemini       GeminiConfig
	Zoom         ZoomConfig
	Drive        DriveConfig
	Notification NotificationConfig
	CORS         CORSConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port string
}

// GeminiConfig holds Gemini API configuration
type GeminiConfig struct {
	APIKey string
}

// ZoomConfig holds Zoom API configuration
type ZoomConfig struct {
	APIKey    string
	APISecret string
}

// DriveConfig holds Google Drive API configuration
type DriveConfig struct {
	APIKey string
}

// NotificationConfig holds notification service configuration
type NotificationConfig struct {
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
		Zoom: ZoomConfig{
			APIKey:    os.Getenv("ZOOM_API_KEY"),
			APISecret: os.Getenv("ZOOM_API_SECRET"),
		},
		Drive: DriveConfig{
			APIKey: os.Getenv("GOOGLE_DRIVE_API_KEY"),
		},
		Notification: NotificationConfig{
			APIKey: os.Getenv("NOTIFICATION_API_KEY"),
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
	// Gemini API Key is optional for basic functionality
	// If not provided, transaction extraction won't work but other features will
	if c.Gemini.APIKey == "" {
		log.Println("⚠️  WARNING: GEMINI_API_KEY not set - transaction extraction will not work")
	}

	// Optional validation for meeting functionality
	if c.Zoom.APIKey == "" {
		// Log warning but don't fail - Zoom functionality won't work
		// In production, you might want to make this required
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
