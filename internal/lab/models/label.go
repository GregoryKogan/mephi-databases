package models

import "gorm.io/gorm"

type Label struct {
	gorm.Model
	BoardID uint   `gorm:"index;not null"`
	Title   string `gorm:"not null"`
	Color   string `gorm:"not null"`
	Cards   []Card `gorm:"many2many:card_labels;"`
}
