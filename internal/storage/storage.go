package storage

import (
	"github.com/gin-gonic/gin"
	"github.com/romainbousseau/probhammer/internal/models"
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

// FindDataSheets returns all datasheets from DB
// TODO: remove, this is a test function 
func (s Storage) FindDatasheets(ctx *gin.Context) ([]*models.Datasheet, error) {
	var datasheets []*models.Datasheet

	err := s.db.Debug().WithContext(ctx).Find(&datasheets).Error
	if err != nil {
		return nil, err
	}

	return datasheets, nil
}
