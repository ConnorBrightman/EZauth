package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Port               string
	Host               string
	JWTSecret          string
	AccessTokenExpiry  time.Duration
	RefreshTokenExpiry time.Duration
	Storage            string
	FilePath           string // used if Storage == file
	LoggingEnabled     bool
}

// LoadConfig loads the configuration from ./ezauth/config.yaml or defaults
func LoadConfig() *Config {
	configDir := "./ezauth"
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)
	viper.AutomaticEnv() // allow overriding via environment variables

	// Default values
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("HOST", "127.0.0.1")
	viper.SetDefault("JWT_SECRET", "super-secret-key")
	viper.SetDefault("ACCESS_TOKEN_EXPIRY", "5m")
	viper.SetDefault("REFRESH_TOKEN_EXPIRY", "168h")
	viper.SetDefault("STORAGE", "memory")
	viper.SetDefault("FILE_PATH", filepath.Join(configDir, "data", "users.json"))
	viper.SetDefault("LOGGING_ENABLED", true)

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No config file found in ./ezauth, using defaults and environment variables")
	}

	// Parse durations
	accessDur, err := time.ParseDuration(viper.GetString("ACCESS_TOKEN_EXPIRY"))
	if err != nil {
		log.Fatalf("Invalid ACCESS_TOKEN_EXPIRY: %v", err)
	}
	refreshDur, err := time.ParseDuration(viper.GetString("REFRESH_TOKEN_EXPIRY"))
	if err != nil {
		log.Fatalf("Invalid REFRESH_TOKEN_EXPIRY: %v", err)
	}

	return &Config{
		Port:               viper.GetString("PORT"),
		Host:               viper.GetString("HOST"),
		JWTSecret:          viper.GetString("JWT_SECRET"),
		AccessTokenExpiry:  accessDur,
		RefreshTokenExpiry: refreshDur,
		Storage:            viper.GetString("STORAGE"),
		FilePath:           viper.GetString("FILE_PATH"),
		LoggingEnabled:     viper.GetBool("LOGGING_ENABLED"),
	}
}

// InitConfig bootstraps the ./ezauth folder with config.yaml and data/users.json
func InitConfig() error {
	configDir := "./ezauth"
	dataDir := filepath.Join(configDir, "data")
	configPath := filepath.Join(configDir, "config.yaml")
	usersPath := filepath.Join(dataDir, "users.json")

	// Don't overwrite existing config
	if _, err := os.Stat(configPath); err == nil {
		return fmt.Errorf("config.yaml already exists at %s", configPath)
	}

	// Create directories
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return fmt.Errorf("failed to create ezauth/data directories: %v", err)
	}

	// Set default values as strings for YAML compatibility
	viper.Set("PORT", "8080")
	viper.Set("HOST", "127.0.0.1")
	viper.Set("JWT_SECRET", "super-secret-key")
	viper.Set("ACCESS_TOKEN_EXPIRY", "5m")
	viper.Set("REFRESH_TOKEN_EXPIRY", "168h")
	viper.Set("STORAGE", "file")
	viper.Set("FILE_PATH", usersPath)
	viper.Set("LOGGING_ENABLED", true)

	// Write config.yaml
	if err := viper.WriteConfigAs(configPath); err != nil {
		return fmt.Errorf("failed to write config.yaml: %v", err)
	}

	// Create empty users.json if it doesn't exist
	if _, err := os.Stat(usersPath); os.IsNotExist(err) {
		if err := os.WriteFile(usersPath, []byte("[]"), 0644); err != nil {
			return fmt.Errorf("failed to create users.json: %v", err)
		}
	}

	log.Println("✅ EZauth initialized successfully in ./ezauth")
	return nil
}
