package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Title       string                 `gorm:"unique;not null"`
	Permissions map[string]interface{} `gorm:"type:jsonb"`
}
