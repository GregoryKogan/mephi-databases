package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	BoardID uint   `gorm:"index;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Title   string `gorm:"not null"`
	Order   int    `gorm:"not null"`
	Cards   []Card `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
