package bucket

import (
	"fmt"

	"github.com/rudransh-shrivastava/self-space/db"
	"gorm.io/gorm"
)

type BucketStore struct {
	db *gorm.DB
}

func NewBucketStore(db *gorm.DB) *BucketStore {
	return &BucketStore{db: db}
}

func (b *BucketStore) CreateBucket(name string) error {
	bucket := db.Bucket{Name: name}
	result := b.db.Create(&bucket)
	if result.Error != nil {
		fmt.Println("bucket with that name already exists")
		return result.Error
	}
	return nil
}

func (b *BucketStore) ListBuckets() ([]db.Bucket, error) {
	var buckets []db.Bucket
	result := b.db.Find(&buckets)
	if result.Error != nil {
		return nil, result.Error
	}
	return buckets, nil
}

func (b *BucketStore) FindBucketByName(name string) (*db.Bucket, error) {
	var bucket db.Bucket
	result := b.db.Where("name = ?", name).First(&bucket)
	if result.Error != nil {
		return nil, result.Error
	}
	return &bucket, nil
}

func (b *BucketStore) DeleteBucket(name string) error {
	var bucket db.Bucket
	result := b.db.Where("name = ?", name).Delete(&bucket)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
