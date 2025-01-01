package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type ApiKey struct {
	ID  uint   `gorm:"primaryKey"`
	Key string `gorm:"unique"`
}

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("apikeys.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&ApiKey{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
