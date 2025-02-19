package main

import (
	"encoding/json"
	_ "embed"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
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
	filter := queryParams.Get("filter")

	var data []map[string]interface{}
	if err := json.Unmarshal(usJson, &data); err != nil {
			http.Error(w, "Failed to parse JSON", http.StatusInternalServerError)
			return
	}

	var filteredData []map[string]interface{}
	for _, item := range data {
			if value, ok := item[filter]; ok {
					filteredData = append(filteredData, map[string]interface{}{filter: value})
			}
	}
	

	response := Response{
		Message: "Hello, World!",
		Data:    usJson[filter],
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}