package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func initConfig() {
	viper.AutomaticEnv()
	viper.AddConfigPath("./config")

	viper.SetDefault("version", "v0.0.0")

	switch os.Getenv("ENV") {
	case "PROD":
		viper.SetConfigName("prod")
	default:
		viper.SetConfigName("dev")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}
