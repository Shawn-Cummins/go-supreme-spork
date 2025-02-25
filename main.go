package main

import (
	_ "embed"
	"log"
	"net/http"
	"supreme-spork/internal/handlers"
)

//go:embed api/us.json
var usJson []byte

func main() {
	http.HandleFunc("/foo", handlers.FooHandler(usJson))
	http.HandleFunc("/geography", handlers.GeoHandler(usJson))
	log.Fatal(http.ListenAndServe(":8080", nil))
}