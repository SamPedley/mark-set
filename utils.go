package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// Version is passed in at build time via the -ldflag flag
var Version string

func initConfig() {
	viper.AddConfigPath("./config")

	viper.SetDefault("version", Version)

	switch os.Getenv("ENV") {
	case "PROD":
		viper.SetConfigName("prod")
	default:
		viper.SetConfigName("dev")
	}

	viper.SetEnvPrefix("app")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}
