package db

import (
	"encoding/json"
	"errors"
	"supreme-spork/internal/models"
	"sync"
)

type InMemoryDB struct {
	data models.Payload
	mu   sync.RWMutex
}

func NewInMemoryDB(jsonData []byte) (*InMemoryDB, error) {
	var payload models.Payload
	if err := json.Unmarshal(jsonData, &payload); err != nil {
		return nil, err
	}
	return &InMemoryDB{
		data: payload,
	}, nil
}

func (db *InMemoryDB) GetGeography() models.Geography {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.data.Geography
}

func (db *InMemoryDB) AddGeographyField(fieldName string, value interface{}) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if db.data.Geography.AdditionalFields == nil {
		db.data.Geography.AdditionalFields = make(map[string]interface{})
	}

	db.data.Geography.AdditionalFields[fieldName] = value
	return nil
}

func (db *InMemoryDB) RemoveGeographyField(fieldName string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if db.data.Geography.AdditionalFields == nil {
		return errors.New("field not found")
	}

	if _, exists := db.data.Geography.AdditionalFields[fieldName]; !exists {
		return errors.New("field not found")
	}

	delete(db.data.Geography.AdditionalFields, fieldName)
	return nil
}
