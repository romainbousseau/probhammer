package storage

import (
	"github.com/gin-gonic/gin"
	"github.com/romainbousseau/probhammer/internal/models"
)

// FindDataSheets returns all datasheets from DB
func (s Storage) FindDatasheets(ctx *gin.Context) ([]*models.Datasheet, error) {
	var datasheets []*models.Datasheet

	err := s.db.WithContext(ctx).Find(&datasheets).Error
	if err != nil {
		return nil, err
	}

	return datasheets, nil
}

// CreateDatasheet creates a new datasheet in DB
func (s Storage) CreateDatasheet(ctx *gin.Context, datasheet *models.Datasheet) error {
	err := s.db.WithContext(ctx).Create(datasheet).Error
	if err != nil {
		return err
	}

	return nil
}

// FindDatasheetByID returns a datasheet from DB
func (s Storage) FindDatasheetByID(ctx *gin.Context, id uint) (*models.Datasheet, error) {
	var datasheet models.Datasheet

	err := s.db.WithContext(ctx).First(&datasheet, id).Error
	if err != nil {
		return nil, err
	}

	return &datasheet, nil
}

// DeleteDatasheet deletes a datasheet from DB
func (s Storage) DeleteDatasheet(ctx *gin.Context, id uint) error {
	err := s.db.WithContext(ctx).Delete(&models.Datasheet{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
