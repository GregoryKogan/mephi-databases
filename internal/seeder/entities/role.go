package entities

import (
	"log/slog"

	"github.com/GregoryKogan/mephi-databases/internal/models"
	"github.com/jackc/pgx/pgtype"
	"gorm.io/gorm"
)

type RoleSeeder interface {
	Seed()
	GetIDs() []uint
}

type RoleSeederImpl struct {
	db  *gorm.DB
	ids []uint
}

func NewRoleSeeder(db *gorm.DB) RoleSeeder {
	return &RoleSeederImpl{db: db}
}

func (s *RoleSeederImpl) Seed() {
	slog.Info("Seeding roles")
	defer slog.Info("Roles seeded")

	var userPermissions pgtype.JSONB
	userPermissions.Set(`{"admin": false}`)

	var adminPermissions pgtype.JSONB
	adminPermissions.Set(`{"admin": true}`)

	roles := []models.Role{
		{
			Title:       "user",
			Permissions: userPermissions,
		},
		{
			Title:       "admin",
			Permissions: adminPermissions,
		},
	}

	for _, role := range roles {
		s.db.Create(&role)
		s.ids = append(s.ids, role.ID)
	}
}

func (s *RoleSeederImpl) GetIDs() []uint {
	return s.ids
}
