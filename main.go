package main

import (
	_ "embed"
	"log"
	"net/http"
	"supreme-spork/internal/db"
	"supreme-spork/internal/handlers"
)

//go:embed api/us.json
var usJson []byte

func main() {
	inMemoryDB, err := db.NewInMemoryDB(usJson)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	http.HandleFunc("/foo", handlers.FooHandler(usJson))
	http.HandleFunc("/geography", handlers.GeoHandler(inMemoryDB))
	http.HandleFunc("/geography/add", handlers.AddGeographyFieldHandler(inMemoryDB))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
