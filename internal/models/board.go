package models

import "gorm.io/gorm"

type Board struct {
	gorm.Model
	OwnerID     uint   `gorm:"index;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Title       string `gorm:"not null"`
	Description string `gorm:"type:text;not null"`
	Members     []BoardMember
	Lists       []List
	Labels      []Label
}
