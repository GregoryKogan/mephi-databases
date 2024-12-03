package entities

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/GregoryKogan/mephi-databases/internal/models"
	"github.com/GregoryKogan/mephi-databases/internal/seeder/selector"
	"github.com/brianvoe/gofakeit/v7"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

type BoardSeeder interface {
	Seed(count uint)
	GetIDs() []uint
	SetUserIDs([]uint)
	GetCount() float64
}

type BoardSeederImpl struct {
	db      *gorm.DB
	ids     []uint
	userIDs []uint
}

func NewBoardSeeder(db *gorm.DB) BoardSeeder {
	return &BoardSeederImpl{db: db}
}

func (s *BoardSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d boards", count))
	defer slog.Info("Boards seeded")

	boards := make([]models.Board, count)
	for i := uint(0); i < count; i++ {
		boards[i] = models.Board{
			OwnerID:     selector.NewSliceSelector().Random(s.userIDs),
			Title:       cases.Title(language.English, cases.Compact).String(gofakeit.Adjective() + " " + gofakeit.Noun()),
			Description: gofakeit.Sentence(10),
			Model: gorm.Model{
				CreatedAt: selector.NewDateSelector().Before(time.Now(), time.Duration(time.Hour*24*30*6)),
			},
		}
	}

	if err := s.db.Create(&boards).Error; err != nil {
		panic(err)
	}

	s.ids = make([]uint, count)
	for i, board := range boards {
		s.ids[i] = board.ID
	}
}

func (s *BoardSeederImpl) GetIDs() []uint {
	return s.ids
}

func (s *BoardSeederImpl) GetCount() float64 {
	return float64(len(s.ids))
}

func (s *BoardSeederImpl) SetUserIDs(ids []uint) {
	s.userIDs = ids
}
