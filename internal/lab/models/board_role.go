package models

import (
	"github.com/jackc/pgx/pgtype"
	"gorm.io/gorm"
)

// BoardRole represents a role that can be assigned to a user for a board.

type BoardRole struct {
	gorm.Model
	Title       string       `gorm:"unique;not null"`
	Permissions pgtype.JSONB `gorm:"type:jsonb;not null"`
}
