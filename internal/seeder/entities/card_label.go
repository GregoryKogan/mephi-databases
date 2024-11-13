package entities

import (
	"fmt"
	"log/slog"

	"github.com/GregoryKogan/mephi-databases/internal/models"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

type CardLabelSeeder interface {
	Seed(count uint)
}

type CardLabelSeederImpl struct {
	db *gorm.DB
}

func NewCardLabelSeeder(db *gorm.DB) CardLabelSeeder {
	return &CardLabelSeederImpl{db: db}
}

func (s *CardLabelSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d card labels", count))

	for created := uint(0); created < count; {
		// pick a random board
		board := models.Board{}
		if err := s.db.Model(&models.Board{}).Order("RANDOM()").Preload("Labels").Preload("Lists.Cards.Labels").First(&board).Error; err != nil {
			continue
		}

		// pick a random label from board.Labels
		if len(board.Labels) == 0 {
			continue
		}
		label := board.Labels[rand.Intn(len(board.Labels))]

		// combine cards from all lists
		var cards []models.Card
		for _, list := range board.Lists {
			cards = append(cards, list.Cards...)
		}

		// pick a random card
		if len(cards) == 0 {
			continue
		}
		card := cards[rand.Intn(len(cards))]

		// check if label is already assigned to card
		alreadyAssigned := false
		for _, l := range card.Labels {
			if l.ID == label.ID {
				alreadyAssigned = true
				break
			}
		}

		if alreadyAssigned {
			continue
		}

		// assign label to card
		card.Labels = append(card.Labels, label)

		if err := s.db.Save(&card).Error; err != nil {
			continue
		}

		created++
	}
}
