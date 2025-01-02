package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Permission string

const (
	READ   Permission = "READ"
	WRITE  Permission = "WRITE"
	DELETE Permission = "DELETE"
)

type APIKey struct {
	ID  uint   `gorm:"primaryKey"`
	Key string `gorm:"unique"`
}

type Bucket struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}

type APIKeyBucketPermission struct {
	ID        uint       `gorm:"primaryKey"`
	APIKeyID  uint       `gorm:"not null"`
	APIKey    APIKey     `gorm:"constraint:OnDelete:CASCADE"`
	BucketID  uint       `gorm:"not null"`
	Bucket    Bucket     `gorm:"constraint:OnDelete:CASCADE"`
	Permision Permission `gorm:"not null"`
}

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&APIKey{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&Bucket{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
