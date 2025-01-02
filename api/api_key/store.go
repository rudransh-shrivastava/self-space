package api

import (
	"github.com/rudransh-shrivastava/self-space/db"
	"gorm.io/gorm"
)

type APIKeyStore struct {
	db *gorm.DB
}

// testing api key
// NYRKv7XgqUnh-FBhuR4U10LT5KW1qdxNtFt-xFOeLRc=
func NewAPIKeyStore(db *gorm.DB) *APIKeyStore {
	return &APIKeyStore{db: db}
}

func (a *APIKeyStore) CreateAPIKey(key string) error {
	apiKey := db.APIKey{Key: key}
	result := a.db.Create(&apiKey)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a *APIKeyStore) FindAPIKeyByKey(key string) (*db.APIKey, error) {
	var apiKey db.APIKey
	result := a.db.Where("key = ?", key).First(&apiKey)
	if result.Error != nil {
		return nil, result.Error
	}
	return &apiKey, nil
}

func (a *APIKeyStore) ListAPIKeys() ([]db.APIKey, error) {
	var apiKeys []db.APIKey
	result := a.db.Find(&apiKeys)
	if result.Error != nil {
		return nil, result.Error
	}
	return apiKeys, nil
}
