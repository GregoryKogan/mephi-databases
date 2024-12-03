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

type LabelSeeder interface {
	Seed(count uint)
	GetRecords() []Record
	SetBoardRecords([]Record)
}

type LabelSeederImpl struct {
	db           *gorm.DB
	records      []Record
	boardRecords []Record
}

func NewLabelSeeder(db *gorm.DB) LabelSeeder {
	return &LabelSeederImpl{db: db}
}

func (s *LabelSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d labels", count))
	defer slog.Info("Labels seeded")

	if len(s.boardRecords) == 0 {
		panic("boardIDs are not set")
	}

	labels := make([]models.Label, count)
	for i := uint(0); i < count; i++ {
		boardRecord := s.boardRecords[selector.NewSliceSelector().Random(len(s.boardRecords))]
		labels[i] = models.Label{
			BoardID: boardRecord.ID,
			Title:   cases.Title(language.English, cases.Compact).String(gofakeit.Noun()),
			Color:   gofakeit.HexColor(),
			Model: gorm.Model{
				CreatedAt: selector.NewDateSelector().BeforeNow(boardRecord.CreatedAt),
			},
		}
	}

	if err := s.db.Create(&labels).Error; err != nil {
		panic(err)
	}

	s.records = make([]Record, count)
	for i, label := range labels {
		s.records[i] = Record{ID: label.ID, CreatedAt: label.CreatedAt}
	}
}

func (s *LabelSeederImpl) GetRecords() []Record {
	return s.records
}

func (s *LabelSeederImpl) SetBoardRecords(records []Record) {
	s.boardRecords = records
}
