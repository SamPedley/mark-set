package main

import (
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
)

// healthCheck returns some basic info about the running application
func healthCheck(w http.ResponseWriter, req *http.Request) {
	res := struct {
		Version string `json:"version"`
	}{viper.GetString("version")}
	json.NewEncoder(w).Encode(res)
}
