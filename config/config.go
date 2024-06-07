package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ServerEnvironment string

const (
	Development ServerEnvironment = "dev"
	Production  ServerEnvironment = "prod"
)

// App config struct
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// Server config struct
type ServerConfig struct {
	AppVersion  string
	Host        string
	Port        string
	Environment ServerEnvironment
}

// MongoDB config struct
type DatabaseConfig struct {
	URI      string
	User     string
	Password string
	Name     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	config := &Config{
		Server: ServerConfig{
			AppVersion:  getEnv("APP_VERSION", "v1"),
			Host:        getEnv("APP_HOST", "0.0.0.0"),
			Port:        getEnv("APP_PORT", "8089"),
			Environment: ServerEnvironment(getEnv("APP_ENV", string(Development))),
		},
		Database: DatabaseConfig{
			URI:  getEnv("DB_URI", "mongodb://localhost:27017"),
			Name: getEnv("DB_NAME", "test"),
		},
	}
	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
