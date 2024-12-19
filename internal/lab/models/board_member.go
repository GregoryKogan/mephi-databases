package models

import "gorm.io/gorm"

// BoardMember represents a user's role in a board.

type BoardMember struct {
	gorm.Model
	UserID      uint `gorm:"index;not null"`
	BoardID     uint `gorm:"index;not null"`
	BoardRoleID uint
	BoardRole   BoardRole `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
