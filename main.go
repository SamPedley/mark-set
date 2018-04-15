package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	initConfig()
}

func main() {
	port := viper.GetString("port")
	errChan := make(chan error, 10)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	r := registerHandlers()

	svr := &http.Server{
		Handler:      r,
		Addr:         port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	go func() {
		errChan <- svr.ListenAndServe()
	}()

	log.Info("listeing on port ", port)
	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case s := <-signalChan:
			log.WithFields(log.Fields{
				"Signal Event": s,
			}).Info("Stopping application")

			err := svr.Shutdown(context.Background())
			if err != nil {
				log.Fatal("Error shuting down server", err)
			}
			os.Exit(0)

		}
	}
}
