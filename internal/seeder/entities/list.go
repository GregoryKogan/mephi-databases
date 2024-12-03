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
	GetRecords() []Record
	GetCount() float64
	SetBoardRecords([]Record)
}

type ListSeederImpl struct {
	db           *gorm.DB
	records      []Record
	boardRecords []Record
}

func NewListSeeder(db *gorm.DB) ListSeeder {
	return &ListSeederImpl{db: db}
}

func (s *ListSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d lists", count))
	defer slog.Info("Lists seeded")

	if len(s.boardRecords) == 0 {
		panic("boardIDs are not set")
	}

	lists := make([]models.List, count)
	for i := uint(0); i < count; i++ {
		boardRecord := s.boardRecords[selector.NewSliceSelector().Random(len(s.boardRecords))]
		lists[i] = models.List{
			BoardID: boardRecord.ID,
			Title:   cases.Title(language.English, cases.Compact).String(gofakeit.Adjective() + " " + gofakeit.Noun()),
			Order:   int(i),
			Model: gorm.Model{
				CreatedAt: selector.NewDateSelector().BeforeNow(boardRecord.CreatedAt),
			},
		}
	}

	if err := s.db.Create(&lists).Error; err != nil {
		panic(err)
	}

	s.records = make([]Record, count)
	for i, list := range lists {
		s.records[i] = Record{ID: list.ID, CreatedAt: list.CreatedAt}
	}
}

func (s *ListSeederImpl) GetRecords() []Record {
	return s.records
}

func (s *ListSeederImpl) GetCount() float64 {
	return float64(len(s.records))
}

func (s *ListSeederImpl) SetBoardRecords(records []Record) {
	s.boardRecords = records
}
