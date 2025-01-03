package apikey

import (
	"errors"
	"fmt"

	"github.com/rudransh-shrivastava/self-space/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type APIKeyStore struct {
	db *gorm.DB
}

// testing api key
// H00fy7Q0_5ErxKL8-Ot48NvjRO82ZQKXSH_4TuJ67NE=
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
	var apiKeys []db.APIKey

	result := a.db.Find(&apiKeys)

	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(apiKeys)
	for _, apiKey := range apiKeys {
		fmt.Println(apiKey.Key)
		err := bcrypt.CompareHashAndPassword([]byte(apiKey.Key), []byte(key))
		if err == nil {
			return &apiKey, nil
		}
	}
	return nil, errors.New("API key not found")
}

func (a *APIKeyStore) ListAPIKeys() ([]db.APIKey, error) {
	var apiKeys []db.APIKey
	result := a.db.Find(&apiKeys)
	if result.Error != nil {
		return nil, result.Error
	}
	return apiKeys, nil
}
