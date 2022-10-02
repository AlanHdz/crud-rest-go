package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port             string
	ConnectionString string
}

func NewConfig(port string, connection string) (*Config, error) {

	if port == "" {
		return nil, errors.New("Port is required")
	}

	if connection == "" {
		return nil, errors.New("Connection is required")
	}

	return &Config{
		Port:             port,
		ConnectionString: connection,
	}, nil
}

func LoadAppConfig() *Config {
	log.Println("Loading server configurations...")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env")
	}

	PORT := os.Getenv("PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=54321 sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)

	config, error := NewConfig(PORT, dsn)

	if error != nil {
		log.Fatal(error)
	}

	return config
}
