package models

import "gorm.io/gorm"

// Datasheet describes a set of rules for a unit
type Datasheet struct {
	gorm.Model

	Name      string `form:"name" json:"name" binding:"required"`
	Toughness int32  `form:"toughness" json:"toughness" binding:"required"`
	Wounds    int32  `form:"wounds" json:"wounds" binding:"required"`
	Save      int32  `form:"save" json:"save" binding:"required"`

	Weapons []WeaponProfile
}

// WeaponProfile describes a set of rules for a weapon
type WeaponProfile struct {
	gorm.Model

	DatasheetID uint

	Name    string
	Attacks int32
	Damage  int32
	Skill   int32
}
