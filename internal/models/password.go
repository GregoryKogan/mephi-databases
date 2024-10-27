package models

import "gorm.io/gorm"

type Password struct {
	gorm.Model
	UserID    uint   `gorm:"index;unique;not null"`
	Hash      string `gorm:"not null"`
	Salt      string `gorm:"unique;not null"`
	Algorithm string `gorm:"not null"`
}
