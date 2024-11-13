package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string `gorm:"not null"`
	Email         string `gorm:"unique;not null"`
	RoleID        uint
	Role          Role `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Password      Password
	Boards        []Board `gorm:"foreignKey:OwnerID"`
	Memberships   []BoardMember
	AssignedCards []Card `gorm:"many2many:card_assignees"`
	Comments      []Comment
}
