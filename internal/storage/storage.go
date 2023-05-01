package storage

import (
	"gorm.io/gorm"
)

// Storage implements database connection with gorm
type Storage struct {
	db *gorm.DB
}

// NewStorage initialize a Storage
func NewStorage(db *gorm.DB) *Storage {
	return &Storage{db}
}
