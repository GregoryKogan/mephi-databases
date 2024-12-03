package entities

import (
	"fmt"
	"log/slog"

	"github.com/GregoryKogan/mephi-databases/internal/models"
	"github.com/GregoryKogan/mephi-databases/internal/seeder/selector"
	"github.com/brianvoe/gofakeit/v7"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

type ListSeeder interface {
	Seed(count uint)
	GetIDs() []uint
	GetCount() float64
	SetBoardIDs([]uint)
}

type ListSeederImpl struct {
	db       *gorm.DB
	ids      []uint
	boardIDs []uint
}

func NewListSeeder(db *gorm.DB) ListSeeder {
	return &ListSeederImpl{db: db}
}

func (s *ListSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d lists", count))
	defer slog.Info("Lists seeded")

	if len(s.boardIDs) == 0 {
		panic("boardIDs are not set")
	}

	lists := make([]models.List, count)
	for i := uint(0); i < count; i++ {
		lists[i] = models.List{
			BoardID: selector.NewSliceSelector().Random(s.boardIDs),
			Title:   cases.Title(language.English, cases.Compact).String(gofakeit.Adjective() + " " + gofakeit.Noun()),
			Order:   int(i),
		}
	}

	if err := s.db.Create(&lists).Error; err != nil {
		panic(err)
	}

	s.ids = make([]uint, count)
	for i, list := range lists {
		s.ids[i] = list.ID
	}
}

func (s *ListSeederImpl) GetIDs() []uint {
	return s.ids
}

func (s *ListSeederImpl) GetCount() float64 {
	return float64(len(s.ids))
}

func (s *ListSeederImpl) SetBoardIDs(ids []uint) {
	s.boardIDs = ids
}
