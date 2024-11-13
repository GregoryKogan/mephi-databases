package entities

import (
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
	var userPermissions pgtype.JSONB
	userPermissions.Set(`{"admin": false}`)

	var adminPermissions pgtype.JSONB
	adminPermissions.Set(`{"admin": true}`)

	roles := []models.Role{
		{
			Title:       "admin",
			Permissions: adminPermissions,
		},
		{
			Title:       "user",
			Permissions: userPermissions,
		},
	}

	// Delete all roles before seeding
	s.db.Exec("DELETE FROM roles")

	for _, role := range roles {
		s.db.Create(&role)
		s.ids = append(s.ids, role.ID)
	}
}

func (s *RoleSeederImpl) GetIDs() []uint {
	return s.ids
}
