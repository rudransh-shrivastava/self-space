package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type APIKey struct {
	ID  uint   `gorm:"primaryKey"`
	Key string `gorm:"unique"`
}

type Bucket struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
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
