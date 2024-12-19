package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	ListID      uint   `gorm:"index;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Title       string `gorm:"not null"`
	Content     string `gorm:"type:text"`
	Order       int    `gorm:"not null"`
	Completed   bool   `gorm:"default:false"`
	DueDate     sql.NullTime
	Attachments []Attachment   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Comments    []Comment      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Assignments []CardAssignee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Labels      []Label        `gorm:"many2many:card_labels;"`
}
