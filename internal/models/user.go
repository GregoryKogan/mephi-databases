package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string `gorm:"unique;not null" faker:"username"`
	Email         string `gorm:"unique;not null" faker:"email"`
	RoleID        uint
	Role          Role `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Password      Password
	Boards        []Board `gorm:"foreignKey:OwnerID"`
	Memberships   []BoardMember
	AssignedCards []Card `gorm:"many2many:card_assignees"`
	Comments      []Comment
}
