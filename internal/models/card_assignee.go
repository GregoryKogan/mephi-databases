package models

import "gorm.io/gorm"

type CardAssignee struct {
	gorm.Model
	UserID uint `gorm:"index;not null"`
	CardID uint `gorm:"index;not null"`
}
