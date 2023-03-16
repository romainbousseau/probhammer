package models

import "gorm.io/gorm"

// Faction describes an faction army from the game
// Subfactions have their ParentID represented as the code of their parent.
type Faction struct {
	gorm.Model

	Name         string `csv:"name"`
	Code         string `csv:"id"`
	ParentID     string `csv:"parent_id"`
	IsSubFaction bool   `csv:"is_subfaction"`
}
