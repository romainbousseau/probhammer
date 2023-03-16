package models

import (
	"time"

	"gorm.io/gorm"
)

// Datasheet describes a set of rules for a datasheet
type Datasheet struct {
	gorm.Model

	ID             uint   `gorm:"primarykey" csv:"id"`
	Name           string `csv:"name"`
	FactionId      string `csv:"faction_id"`
	DatasheetUnits []Unit

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
