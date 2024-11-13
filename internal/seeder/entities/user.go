package entities

import (
	"log/slog"

	"github.com/GregoryKogan/mephi-databases/internal/models"
	"github.com/go-faker/faker/v4"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

type UserSeeder interface {
	Seed(count uint)
	SetRoleIDs(roleIDs []uint)
	GetIDs() []uint
}

type UserSeederImpl struct {
	db      *gorm.DB
	roleIDs []uint
	ids     []uint
}

func NewUserSeeder(db *gorm.DB) UserSeeder {
	return &UserSeederImpl{db: db}
}

func (s *UserSeederImpl) Seed(count uint) {
	slog.Info("Seeding users")

	if len(s.roleIDs) == 0 {
		panic("roleIDs are not set")
	}

	for i := uint(0); i < count; i++ {
		user := models.User{
			Username: faker.FirstName() + " " + faker.LastName(),
			Email:    faker.Email(),
			RoleID:   s.roleIDs[rand.Intn(len(s.roleIDs))],
		}

		s.db.Create(&user)
		s.ids = append(s.ids, user.ID)
	}
}

func (s *UserSeederImpl) SetRoleIDs(roleIDs []uint) {
	s.roleIDs = roleIDs
}

func (s *UserSeederImpl) GetIDs() []uint {
	return s.ids
}
