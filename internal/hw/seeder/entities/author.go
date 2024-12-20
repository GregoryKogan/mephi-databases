package entities

import (
	"fmt"
	"log/slog"

	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

type AuthorSeeder interface {
	Seed(count uint)
	GetIDs() []uint
}

type Author struct {
	AId   uint   `gorm:"primaryKey;column:a_id"`
	AName string `gorm:"column:a_name"`
}

type AuthorSeederImpl struct {
	db  *gorm.DB
	IDs []uint
}

func NewAuthorSeeder(db *gorm.DB) AuthorSeeder {
	return &AuthorSeederImpl{db: db}
}

func (s *AuthorSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d authors", count))
	defer slog.Info("Authors seeded")

	authors := make([]Author, count)
	for i := uint(0); i < count; i++ {
		name := gofakeit.Name() + " " + gofakeit.LastName()
		author := Author{
			AName: name,
		}
		authors[i] = author
	}

	if err := s.db.Create(&authors).Error; err != nil {
		panic(err)
	}

	s.IDs = make([]uint, count)
	for i, author := range authors {
		s.IDs[i] = author.AId
	}
}

func (s *AuthorSeederImpl) GetIDs() []uint {
	return s.IDs
}
