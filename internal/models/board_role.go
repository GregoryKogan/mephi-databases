package models

import "gorm.io/gorm"

// BoardRole represents a role that can be assigned to a user for a board.

type BoardRole struct {
	gorm.Model
	Title       string                 `gorm:"unique;not null"`
	Permissions map[string]interface{} `gorm:"type:jsonb"`
}
