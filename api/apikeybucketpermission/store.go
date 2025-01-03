package apikeybucketpermission

import (
	"github.com/rudransh-shrivastava/self-space/db"
	"gorm.io/gorm"
)

type APIKeyBucketPermissionStore struct {
	db *gorm.DB
}

func NewAPIKeyBucketPermissionStore(db *gorm.DB) *APIKeyBucketPermissionStore {
	return &APIKeyBucketPermissionStore{db: db}
}

func (a *APIKeyBucketPermissionStore) CreateAPIKeyBucketPermission(apiKeyID, bucketID uint, permission string) error {
	apiKeyBucketPermission := db.APIKeyBucketPermission{
		APIKeyID:  apiKeyID,
		BucketID:  bucketID,
		Permision: db.Permission(permission),
	}
	result := a.db.Create(&apiKeyBucketPermission)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a *APIKeyBucketPermissionStore) HasPermission(apiKey *db.APIKey, bucket *db.Bucket, permission string) (bool, error) {
	return true, nil
}
