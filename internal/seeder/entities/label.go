package entities

import (
	"log/slog"
	"math/rand"

	"github.com/GregoryKogan/mephi-databases/internal/models"
	"github.com/go-faker/faker/v4"
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
	slog.Info("Seeding labels")

	if len(s.boardIDs) == 0 {
		panic("boardIDs are not set")
	}

	for i := uint(0); i < count; i++ {
		label := models.Label{
			BoardID: s.boardIDs[rand.Intn(len(s.boardIDs))],
			Title:   "Label " + faker.Word(),
			Color:   randomHexColor(),
		}

		s.db.Create(&label)
		s.ids = append(s.ids, label.ID)
	}
}

func (s *LabelSeederImpl) GetIDs() []uint {
	return s.ids
}

func (s *LabelSeederImpl) SetBoardIDs(ids []uint) {
	s.boardIDs = ids
}

func randomHexColor() string {
	chars := "abcdef0123456789"
	color := "#"
	for i := 0; i < 6; i++ {
		color += string(chars[rand.Intn(len(chars))])
	}
	return color
}
