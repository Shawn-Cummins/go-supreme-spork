package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"supreme-spork/internal/models"
)

func FooHandler(p []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	log.Println(queryParams)

	var payload models.Payload
	if err := json.Unmarshal(p, &payload); err != nil {
		http.Error(w, "Failed to parse JSON result test", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	totalAreaText := payload.Geography.Coastline.Text
    fmt.Println("Geography.Area.total.text:", totalAreaText)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload.People.DependencyRatios)
	}
}