package api

import (
	"gorm.io/gorm"
)

type BucketStore struct {
	db *gorm.DB
}

func NewBucketStore(db *gorm.DB) *BucketStore {
	return &BucketStore{db: db}
}

func (b *BucketStore) CreateBucket(name string) error {
	// create bucket here
	return nil
}
