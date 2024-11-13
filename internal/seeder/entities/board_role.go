package entities

import (
	"github.com/GregoryKogan/mephi-databases/internal/models"
	"github.com/jackc/pgx/pgtype"
	"gorm.io/gorm"
)

type BoardRoleSeeder interface {
	Seed()
	GetIDs() []uint
}

type BoardRoleSeederImpl struct {
	db  *gorm.DB
	ids []uint
}

func NewBoardRoleSeeder(db *gorm.DB) BoardRoleSeeder {
	return &BoardRoleSeederImpl{db: db}
}

func (s *BoardRoleSeederImpl) Seed() {
	var roles = []models.BoardRole{
		{
			Title: "redactor",
			Permissions: pgtype.JSONB{
				Bytes:  []byte(`{"read": true, "write": true, "delete": true}`),
				Status: pgtype.Present,
			},
		},
		{
			Title: "commentator",
			Permissions: pgtype.JSONB{
				Bytes:  []byte(`{"read": true, "write": true, "delete": false}`),
				Status: pgtype.Present,
			},
		},
		{
			Title: "viewer",
			Permissions: pgtype.JSONB{
				Bytes:  []byte(`{"read": true, "write": false, "delete": false}`),
				Status: pgtype.Present,
			},
		},
	}

	// Delete all roles before seeding
	s.db.Exec("DELETE FROM board_roles")

	for _, role := range roles {
		s.db.Create(&role)
		s.ids = append(s.ids, role.ID)
	}
}

func (s *BoardRoleSeederImpl) GetIDs() []uint {
	return s.ids
}
