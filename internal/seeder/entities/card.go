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

	if len(s.listIDs) == 0 {
		panic("listIDs are not set")
	}

	for i := uint(0); i < count; i++ {
		card := models.Card{
			ListID:  s.listIDs[rand.Intn(len(s.listIDs))],
			Title:   "Card " + faker.Word(),
			Content: faker.Sentence(),
			Order:   int(i),
		}

		s.db.Create(&card)
		s.ids = append(s.ids, card.ID)
	}
}

func (s *CardSeederImpl) GetIDs() []uint {
	return s.ids
}

func (s *CardSeederImpl) SetListIDs(ids []uint) {
	s.listIDs = ids
}
