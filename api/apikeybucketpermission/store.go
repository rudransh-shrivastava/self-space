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
		APIKeyID:   apiKeyID,
		BucketID:   bucketID,
		Permission: db.Permission(permission),
	}
	result := a.db.Create(&apiKeyBucketPermission)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a *APIKeyBucketPermissionStore) HasPermission(apiKey *db.APIKey, bucket *db.Bucket, permission string) (bool, error) {
	var apiKeyBucketPermission db.APIKeyBucketPermission
	// find apikeybucketpermisssion where APIKey = apikey and Bucket = bucket and Pemission = permission
	result := a.db.Where("api_key_id = ? AND bucket_id = ? AND permission = ?", apiKey.ID, bucket.ID, db.Permission(permission)).First(&apiKeyBucketPermission)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func (a *APIKeyBucketPermissionStore) DeletePermission(apiKey *db.APIKey, bucket *db.Bucket, permission string) error {
	var apiKeyBucketPermission db.APIKeyBucketPermission

	result := a.db.Where("api_key_id = ? AND bucket_id = ? AND permission = ?", apiKey.ID, bucket.ID, db.Permission(permission)).Delete(&apiKeyBucketPermission)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
