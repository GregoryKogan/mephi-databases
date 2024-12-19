package models

import (
	"github.com/jackc/pgx/pgtype"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Title       string       `gorm:"unique;not null"`
	Permissions pgtype.JSONB `gorm:"type:jsonb;not null"`
}
