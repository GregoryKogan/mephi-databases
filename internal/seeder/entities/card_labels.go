package entities

import (
	"fmt"
	"log/slog"

	"github.com/GregoryKogan/mephi-databases/internal/models"
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
		// pick a random card
		card := models.Card{}
		if err := s.db.Model(&models.Card{}).Order("RANDOM()").First(&card).Error; err != nil {
			continue
		}

		list := models.List{}
		if err := s.db.First(&list, card.ListID).Error; err != nil {
			continue
		}

		// pick a random label from same board
		label := models.Label{}
		if err := s.db.Model(&models.Label{}).Where("board_id = ?", list.BoardID).Order("RANDOM()").First(&label).Error; err != nil {
			continue
		}

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
		if err := s.db.Model(&card).Association("Labels").Append(&label); err != nil {
			continue
		}

		created++
	}
}
