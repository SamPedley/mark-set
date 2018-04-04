package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func mainRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(withContext)
	r.Use(withLogging)

	r.HandleFunc("/", getRoot).Methods("GET")
	r.HandleFunc("/info", getInfo).Methods("GET")

	return r
}

func getRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

// getInfo returns some basic info about the running application
func getInfo(w http.ResponseWriter, req *http.Request) {
	res := struct {
		Version string `json:"version"`
	}{viper.GetString("version")}
	json.NewEncoder(w).Encode(res)
}

type requestIDKey string

var requestID = requestIDKey("requestID")

// withContext is a middleware that adds context to a request with a unique requestID
func withContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestID, id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// withLogging logs the inital request and the response
func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ctx := r.Context()
		id, ok := ctx.Value(requestID).(int64)
		if !ok {
			id = 0
		}

		log.WithFields(log.Fields{
			"requestID": id,
			"method":    r.Method,
			"url":       r.URL,
		}).Info("Request Recived")

		go func() {
			select {
			case <-ctx.Done():
				log.WithFields(log.Fields{
					"requestID": id,
					"elapsed":   time.Now().Sub(start),
				}).Info("Request Finished")
			}
		}()

		next.ServeHTTP(w, r)
	})
}
