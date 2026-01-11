package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type Config struct{
	DBHost     string
	DBName     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBMode     string
	DBTimezone string
}

func Init() (*Config, error){
	error := godotenv.Load()

	if error != nil {
		return nil, fmt.Errorf("Error loading .env file")
	}

	config := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_DATABASE"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBMode:     os.Getenv("DB_SSLMODE"),
		DBTimezone: os.Getenv("DB_TIMEZONE"),
	}

	return config, nil
}