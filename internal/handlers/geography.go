package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"supreme-spork/internal/models"
)

func GeoHandler(p []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	var payload models.Payload
	if err := json.Unmarshal(p, &payload); err != nil {
		http.Error(w, "Failed to parse JSON result test", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload.Geography)
	}
}