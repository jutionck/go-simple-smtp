package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	ApiPort string
	ApiHost string
}

type EmailConfig struct {
	Server    string
	Port      string
	EmailFrom string
	Password  string
}

type Config struct {
	ApiConfig
	EmailConfig
}

func (c *Config) readConfigFile() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c.EmailConfig = EmailConfig{
		Server:    os.Getenv("EMAIL_SERVER"),
		Port:      os.Getenv("EMAIL_PORT"),
		EmailFrom: os.Getenv("EMAIL_FROM"),
		Password:  os.Getenv("EMAIL_PASSWORD"),
	}

	c.ApiConfig = ApiConfig{
		ApiHost: os.Getenv("API_HOST"),
		ApiPort: os.Getenv("API_PORT"),
	}

	if c.EmailConfig.Server == "" || c.EmailConfig.Port == "" || c.EmailConfig.EmailFrom == "" ||
		c.EmailConfig.Password == "" || c.ApiConfig.ApiHost == "" || c.ApiConfig.ApiPort == "" {
		return errors.New("missing required environment variables")
	}
	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.readConfigFile()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
