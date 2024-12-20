package entities

import (
	"fmt"
	"log/slog"

	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

type BookSeeder interface {
	Seed(count uint)
	GetIDs() []uint
}

type Book struct {
	BId       uint   `gorm:"primaryKey;column:b_id"`
	BName     string `gorm:"column:b_name"`
	BYear     uint   `gorm:"column:b_year"`
	BQuantity uint   `gorm:"column:b_quantity"`
}

type BookSeederImpl struct {
	db  *gorm.DB
	IDs []uint
}

func NewBookSeeder(db *gorm.DB) BookSeeder {
	return &BookSeederImpl{db: db}
}

func (s *BookSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d Books", count))
	defer slog.Info("Books seeded")

	Books := make([]Book, count)
	for i := uint(0); i < count; i++ {
		Book := Book{
			BName:     gofakeit.BookTitle(),
			BYear:     uint(gofakeit.Year()),
			BQuantity: uint(gofakeit.Number(1, 100)),
		}
		Books[i] = Book
	}

	if err := s.db.Create(&Books).Error; err != nil {
		panic(err)
	}

	s.IDs = make([]uint, count)
	for i, Book := range Books {
		s.IDs[i] = Book.BId
	}
}

func (s *BookSeederImpl) GetIDs() []uint {
	return s.IDs
}
