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
	DueDate     sql.NullTime
	Attachments []Attachment
	Assignees   []User `gorm:"many2many:card_assignees;"`
	Comments    []Comment
	Labels      []Label `gorm:"many2many:card_labels;"`
}
