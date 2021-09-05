package handler

import (
	"encoding/json"
	"net/http"
)

type version struct {
	Version string `json:"version"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	response := version{Version: "0.0.0"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
