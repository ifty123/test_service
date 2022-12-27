package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

var EnvConfigs *envConfigs

// struct to map env values
type envConfigs struct {
	LocalServerPort string `mapstructure:"SERVER_PORT"`
	DbUsername      string `mapstructure:"DB_USERNAME"`
	DbPassword      string `mapstructure:"DB_PASSWORD"`
	DbName          string `mapstructure:"DB_NAME"`
	DbHost          string `mapstructure:"DB_HOST"`
	DbPort          string `mapstructure:"DB_PORT"`
	AppName         string `mapstructure:"APP_NAME"`
}

func loadEnvVariables() (config *envConfigs) {
	// Tell viper the path/location of your env file. If it is root just add "."
	viper.AddConfigPath(".")

	// Tell viper the name of your file
	viper.SetConfigName(".env")

	// Tell viper the type of your file
	viper.SetConfigType("env")

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
