package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"context"
	"supreme-spork/internal/db"
	"supreme-spork/internal/models"
)

type AddFieldRequest struct {
	FieldName string      `json:"field_name"`
	Value     interface{} `json:"value"`
}

type ServerImpl struct {
	DB CountryDB
}

type CountryDB interface {
	GetGeography(context.Context) *models.Geography
	UpsertGeography(context.Context, models.Geography)
	GetIntro(context.Context) *models.Introduction
}

func (s *ServerImpl) GetGeography(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetGeography called")
}

func (s *ServerImpl) PostGeographyAdd(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostGeographyAdd called")
}

func (s *ServerImpl) GeoHandler(db *db.InMemoryDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GeoHandler called")

		geography := db.GetGeography(r.Context())

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(geography); err != nil {
			http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
			fmt.Println("Error encoding JSON response:", err)
		}
	}
}

func (s *ServerImpl) GetIntro(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetIntro called")
}