package entities

import (
	"fmt"
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
	slog.Info(fmt.Sprintf("Seeding %d users", count))

	if len(s.roleIDs) == 0 {
		panic("roleIDs are not set")
	}

	users := make([]models.User, count)
	for i := uint(0); i < count; i++ {
		users[i] = models.User{
			Username: faker.FirstName() + " " + faker.LastName(),
			Email:    faker.Email(),
			RoleID:   s.roleIDs[rand.Intn(len(s.roleIDs))],
		}
	}

	if err := s.db.Create(&users).Error; err != nil {
		panic(err)
	}

	s.ids = make([]uint, count)
	for i, user := range users {
		s.ids[i] = user.ID
	}
}

func (s *UserSeederImpl) SetRoleIDs(roleIDs []uint) {
	s.roleIDs = roleIDs
}

func (s *UserSeederImpl) GetIDs() []uint {
	return s.ids
}
