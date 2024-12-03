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
	GetIDs() []uint
	SetBoardIDs([]uint)
}

type LabelSeederImpl struct {
	db       *gorm.DB
	ids      []uint
	boardIDs []uint
}

func NewLabelSeeder(db *gorm.DB) LabelSeeder {
	return &LabelSeederImpl{db: db}
}

func (s *LabelSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d labels", count))
	defer slog.Info("Labels seeded")

	if len(s.boardIDs) == 0 {
		panic("boardIDs are not set")
	}

	labels := make([]models.Label, count)
	for i := uint(0); i < count; i++ {
		labels[i] = models.Label{
			BoardID: selector.NewSliceSelector().Random(s.boardIDs),
			Title:   cases.Title(language.English, cases.Compact).String(gofakeit.Noun()),
			Color:   gofakeit.HexColor(),
		}
	}

	if err := s.db.Create(&labels).Error; err != nil {
		panic(err)
	}

	s.ids = make([]uint, count)
	for i, label := range labels {
		s.ids[i] = label.ID
	}
}

func (s *LabelSeederImpl) GetIDs() []uint {
	return s.ids
}

func (s *LabelSeederImpl) SetBoardIDs(ids []uint) {
	s.boardIDs = ids
}
