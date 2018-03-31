package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	initConfig()
}

func main() {
	port := viper.GetString("port")
	log.Info("listeing on port ", port)

	http.HandleFunc("/", healthCheck)
	http.ListenAndServe(port, nil)
}
