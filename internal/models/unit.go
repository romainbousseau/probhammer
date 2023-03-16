package models

import "gorm.io/gorm"

// Unit describes a unit and  its stats in game
type Unit struct {
	gorm.Model

	DatasheetID    uint   `csv:"datasheet_id"`
	Position       int32  `csv:"line"`
	Name           string `csv:"name"`
	WeaponSkill    int32  `csv:"WS"`
	BallisticSkill int32  `csv:"BS"`
	Strength       int32  `csv:"S"`
	Toughness      int32  `csv:"T"`
	Wounds         int32  `csv:"W"`
	// Attacks are string types cause they can have "Dn+n" forms
	Attacks          string `csv:"A"`
	Save             int32  `csv:"Sv"`
	InvulnerableSave int32
}
