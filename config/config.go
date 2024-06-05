package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type ServerEnvironment string

const (
	DEVELOPMENT ServerEnvironment = "dev"
	PRODUCTION  ServerEnvironment = "prod"

	SERVER_APP_VERSION = "v1"
	SERVER_PORT        = "8089"

	LOGGER_DEVELOPMENT = true
	LOGGER_LOG_LEVEL   = "debug"

	DB_HOST     = ""
	DB_PORT     = ""
	DB_USER     = ""
	DB_PASSWORD = ""
	DB_NAME     = ""
	DB_DRIVER   = "postgres"
)

// App config struct
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Metrics  Metrics
	Logger   Logger
}

// Server config struct
type ServerConfig struct {
	AppVersion  string
	Host        string
	Port        string
	Environment ServerEnvironment
}

// Logger config
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// Postgresql config
type DatabaseConfig struct {
	Driver       string
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
	SSLMode      bool
}

// Metrics config
type Metrics struct {
	URL         string
	ServiceName string
}

// Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
