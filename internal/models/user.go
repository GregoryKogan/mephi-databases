package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"not null"`
	Email       string `gorm:"unique;not null"`
	RoleID      uint
	Role        Role           `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Password    Password       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Boards      []Board        `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Memberships []BoardMember  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Comments    []Comment      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Assignments []CardAssignee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
