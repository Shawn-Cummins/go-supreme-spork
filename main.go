package main

import (
	_ "embed"
	"log"
	"net/http"
	"supreme-spork/internal/db"
	"supreme-spork/internal/handlers"
	"supreme-spork/internal/api"
)

//go:embed api/us.json
var usJson []byte

// func main() {
// 	http.Handle("/api", api.HandlerWithOptions(&handlers.ServerImpl{}, api.StdHTTPServerOptions{}))
// 	// http.HandleFunc("/foo", handlers.FooHandler(usJson))
// 	// http.HandleFunc("/geography", handlers.GeoHandler(inMemoryDB))
// 	// http.HandleFunc("/geography/add", handlers.AddGeographyFieldHandler(inMemoryDB))
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

func main() {
	inMemoryDB, err := db.NewInMemoryDB(usJson)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	// Create an instance of your implementation.
	var server api.ServerInterface = &handlers.ServerImpl{
		DB: inMemoryDB,
	}
 
	// Create a new HTTP multiplexer.
	mux := http.NewServeMux()
 
	// Map your implementation to the generated routes.
	// The function name may differ (e.g., NewRouter or HandlerFromMux) depending on your oapi-codegen settings.
	api.HandlerFromMux(server, mux)
 
	// Start the server.
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
 }