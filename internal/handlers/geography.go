package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"supreme-spork/internal/db"
)

type AddFieldRequest struct {
	FieldName string      `json:"field_name"`
	Value     interface{} `json:"value"`
}

type ServerImpl struct {
	FieldName string      `json:"field_name"`
	Value     interface{} `json:"value"`
}

func (s *ServerImpl) GetFoo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetFoo called")
}

func (s *ServerImpl) GetGeography(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetGeography called")
}
func (s *ServerImpl) PostGeographyAdd(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostGeographyAdd called")
}

func GeoHandler(db *db.InMemoryDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GeoHandler called")

		geography := db.GetGeography()

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(geography); err != nil {
			http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
			fmt.Println("Error encoding JSON response:", err)
		}
	}
}

func AddGeographyFieldHandler(db *db.InMemoryDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var req AddFieldRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			fmt.Println("Error decoding request body:", err)
			return
		}

		if err := db.AddGeographyField(req.FieldName, req.Value); err != nil {
			http.Error(w, "Failed to add field", http.StatusInternalServerError)
			fmt.Println("Error adding field:", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Field added successfully")
	}
}
