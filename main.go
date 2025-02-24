package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"supreme-spork/models"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Config struct {
	Port int    `json:"port"`
	Host string `json:"host"`
}

// GenericMap is a map with string keys and values of any type T.
type GenericMap[T any] map[string]T

//go:embed api/us.json
var usJson []byte


func main() {
	http.HandleFunc("/foo", fooHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	log.Println(queryParams)

	var payload models.Payload
	if err := json.Unmarshal(usJson, &payload); err != nil {
		http.Error(w, "Failed to parse JSON result test", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	totalAreaText := payload.Geography.Coastline.Text
    fmt.Println("Geography.Area.total.text:", totalAreaText)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload.People.DependencyRatios)
}
