package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

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
