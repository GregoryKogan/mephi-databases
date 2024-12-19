package models

import "gorm.io/gorm"

type Attachment struct {
	gorm.Model
	CardID  uint
	FileURL string `gorm:"not null"`
}
