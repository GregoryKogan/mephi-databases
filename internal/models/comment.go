package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	CardID uint   `gorm:"index;not null"`
	UserID uint   `gorm:"index;not null"`
	Text   string `gorm:"type:text;not null"`
}
