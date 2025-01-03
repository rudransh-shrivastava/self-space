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
	ID         uint       `gorm:"primaryKey"`
	APIKeyID   uint       `gorm:"not null;foreignKey:APIKeyID;constraint:OnDelete:CASCADE"`
	APIKey     APIKey     `gorm:"constraint:OnDelete:CASCADE"`
	BucketID   uint       `gorm:"not null;foreignKey:BucketID;constraint:OnDelete:CASCADE"`
	Bucket     Bucket     `gorm:"constraint:OnDelete:CASCADE"`
	Permission Permission `gorm:"not null"`
}

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	sqlDB.Exec("PRAGMA foreign_keys = ON")

	err = db.AutoMigrate(&APIKey{}, &Bucket{}, &APIKeyBucketPermission{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
