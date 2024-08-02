package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
)

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	DBHost = viper.GetString("DB_HOST")
	DBPort = viper.GetString("DB_PORT")
	DBUser = viper.GetString("DB_USER")
	DBPassword = viper.GetString("DB_PASSWORD")
	DBName = viper.GetString("DB_NAME")
}
