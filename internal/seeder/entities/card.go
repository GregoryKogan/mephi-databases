package entities

import (
	"fmt"
	"log/slog"

	"github.com/GregoryKogan/mephi-databases/internal/models"
	"github.com/go-faker/faker/v4"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

type CardSeeder interface {
	Seed(count uint)
	GetIDs() []uint
	GetCount() float64
	SetListIDs([]uint)
}

type CardSeederImpl struct {
	db      *gorm.DB
	ids     []uint
	listIDs []uint
}

func NewCardSeeder(db *gorm.DB) CardSeeder {
	return &CardSeederImpl{db: db}
}

func (s *CardSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d cards", count))
	defer slog.Info("Cards seeded")

	if len(s.listIDs) == 0 {
		panic("listIDs are not set")
	}

	cards := make([]models.Card, count)
	for i := uint(0); i < count; i++ {
		cards[i] = models.Card{
			ListID:  s.listIDs[rand.Intn(len(s.listIDs))],
			Title:   "Card " + faker.Word(),
			Content: faker.Sentence(),
			Order:   int(i),
		}
	}

	if err := s.db.Create(&cards).Error; err != nil {
		panic(err)
	}

	s.ids = make([]uint, count)
	for i, card := range cards {
		s.ids[i] = card.ID
	}
}

func (s *CardSeederImpl) GetIDs() []uint {
	return s.ids
}

func (s *CardSeederImpl) GetCount() float64 {
	return float64(len(s.ids))
}

func (s *CardSeederImpl) SetListIDs(ids []uint) {
	s.listIDs = ids
}
