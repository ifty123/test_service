package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// struct to map env values
type EnvConfigs struct {
	LocalServerPort string
	DbUsername      string
	DbPassword      string
	DbName          string
	DbHost          string
	DbPort          string
	AppName         string
}

func Config() *EnvConfigs {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file", err)
		return nil
	}
	return &EnvConfigs{
		LocalServerPort: os.Getenv("SERVER_PORT"),
		DbUsername:      os.Getenv("DB_USER"),
		DbPassword:      os.Getenv("DB_PASS"),
		DbName:          os.Getenv("DB_NAME"),
		DbHost:          os.Getenv("DB_HOST"),
		DbPort:          os.Getenv("DB_PORT"),
		AppName:         os.Getenv("APP_NAME"),
	}
	// return &EnvConfigs{
	// 	LocalServerPort: "3030",
	// 	DbUsername:      "root",
	// 	DbPassword:      "",
	// 	DbName:          "dbmaster",
	// 	DbHost:          "database",
	// 	DbPort:          "3306",
	// 	AppName:         "test_service",
	// }
}
