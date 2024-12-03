package entities

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/GregoryKogan/mephi-databases/internal/models"
	"github.com/GregoryKogan/mephi-databases/internal/seeder/selector"
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

type UserSeeder interface {
	Seed(count uint)
	SetRoleIDs([]uint)
	GetRecords() []Record
}

type UserSeederImpl struct {
	db      *gorm.DB
	roleIDs []uint
	records []Record
}

func NewUserSeeder(db *gorm.DB) UserSeeder {
	return &UserSeederImpl{db: db}
}

func (s *UserSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d users", count))
	defer slog.Info("Users seeded")

	if len(s.roleIDs) == 0 {
		panic("roleIDs are not set")
	}

	users := make([]models.User, count)
	for i := uint(0); i < count; i++ {
		users[i] = models.User{
			Username: gofakeit.Username(),
			Email:    gofakeit.Email(),
			RoleID:   s.roleIDs[selector.NewSliceSelector().Exponential(len(s.roleIDs))],
			Model: gorm.Model{
				CreatedAt: selector.NewDateSelector().Before(time.Now(), time.Duration(time.Hour*24*30*12)),
			},
		}
	}

	if err := s.db.Create(&users).Error; err != nil {
		panic(err)
	}

	s.records = make([]Record, count)
	for i, user := range users {
		s.records[i] = Record{ID: user.ID, CreatedAt: user.CreatedAt}
	}
}

func (s *UserSeederImpl) SetRoleIDs(roleIDs []uint) {
	s.roleIDs = roleIDs
}

func (s *UserSeederImpl) GetRecords() []Record {
	return s.records
}
