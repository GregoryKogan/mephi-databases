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

type BoardSeeder interface {
	Seed(count uint)
	GetRecords() []Record
	SetUserRecords([]Record)
	GetCount() float64
}

type BoardSeederImpl struct {
	db          *gorm.DB
	records     []Record
	userRecords []Record
}

func NewBoardSeeder(db *gorm.DB) BoardSeeder {
	return &BoardSeederImpl{db: db}
}

func (s *BoardSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d boards", count))
	defer slog.Info("Boards seeded")

	boards := make([]models.Board, count)
	for i := uint(0); i < count; i++ {
		ownerRecord := s.userRecords[selector.NewSliceSelector().Random(len(s.userRecords))]
		boards[i] = models.Board{
			OwnerID:     ownerRecord.ID,
			Title:       cases.Title(language.English, cases.Compact).String(gofakeit.Adjective() + " " + gofakeit.Noun()),
			Description: gofakeit.Sentence(10),
			Model: gorm.Model{
				CreatedAt: selector.NewDateSelector().BeforeNow(ownerRecord.CreatedAt),
			},
		}
	}

	if err := s.db.Create(&boards).Error; err != nil {
		panic(err)
	}

	s.records = make([]Record, count)
	for i, board := range boards {
		s.records[i] = Record{ID: board.ID, CreatedAt: board.CreatedAt}
	}
}

func (s *BoardSeederImpl) GetRecords() []Record {
	return s.records
}

func (s *BoardSeederImpl) GetCount() float64 {
	return float64(len(s.records))
}

func (s *BoardSeederImpl) SetUserRecords(records []Record) {
	s.userRecords = records
}
