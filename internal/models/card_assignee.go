package models

import "gorm.io/gorm"

type CardAssignee struct {
	gorm.Model
	UserID uint `gorm:"index;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CardID uint `gorm:"index;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
