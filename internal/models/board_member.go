package models

import "gorm.io/gorm"

// BoardMember represents a user's role in a board.

type BoardMember struct {
	gorm.Model
	UserID      uint `gorm:"index;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BoardID     uint `gorm:"index;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BoardRoleID uint
	BoardRole   BoardRole `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
