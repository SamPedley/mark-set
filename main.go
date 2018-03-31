package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	initConfig()
}

func main() {
	port := viper.GetString("port")

	r := mux.NewRouter()
	r.HandleFunc("/", getRoot).Methods("GET")
	r.HandleFunc("/info", getInfo).Methods("GET")

	svr := &http.Server{
		Handler:      r,
		Addr:         port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Info("listeing on port ", port)
	log.Fatal(svr.ListenAndServe())
}
