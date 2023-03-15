package models

import "gorm.io/gorm"

// Datasheet describes a set of rules for a unit
type Datasheet struct {
	gorm.Model

	Name string
}
